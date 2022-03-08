package iptables

import (
	"errors"
	"fmt"
	"net"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type Table string

type Chain string

//iptables-extensions manual page
type Target string

type Action string

//iptables entry
type IPTable struct{}

type TableInfo struct {
	Name    Table
	Chains  []Chain
	Targets []Target
}

// content of rule for iptables command
type RuleInfo struct {
	Table  *TableInfo
	Chain  Chain
	Action Action
	Args   []string
}

type RuleError struct {
	Cmd    string
	Output []byte
}

func (e RuleError) Error() string {
	return fmt.Sprintf("Error iptables cmd: %s \noutput: %s", e.Cmd, string(e.Output))
}

const (
	Nat    Table = "nat"
	Filter Table = "filter"
	Mangle Table = "mangle"

	Input       Chain = "INPUT"
	Output      Chain = "OUTPUT"
	Forward     Chain = "FORWARD"
	PreRouting  Chain = "PREROUTING"
	PostRouting Chain = "POSTROUTING"

	Accept     Target = "ACCEPT"
	Drop       Target = "DROP"
	DNat       Target = "DNAT"
	SNat       Target = "SNAT"
	Masquerade Target = "MASQUERADE"

	Append      Action = "-A"
	Delete      Action = "-D"
	Insert      Action = "-I"
	List        Action = "-L"
	NewChain    Action = "-N"
	DeleteChain Action = "-X"
	PolicySet   Action = "-P"
	Save        Action = "-S" // -S --list-rules, printed like iptables-save
)

var (
	NatTable TableInfo = TableInfo{
		Name:    Nat,
		Chains:  []Chain{Output, PreRouting, PreRouting},
		Targets: []Target{DNat, SNat, Masquerade},
	}
	FilterTable TableInfo = TableInfo{
		Name:    Filter,
		Chains:  []Chain{Forward, Input, Output},
		Targets: []Target{Accept, Drop},
	}
	MangleTable TableInfo = TableInfo{
		Name:   Mangle,
		Chains: []Chain{Forward, Input, Output, PostRouting, PreRouting},
	}
	IptablesPath             string
	supportXlock             = false
	statefulAction           = []Action{Append, Insert, Delete}
	ErrIptablesNotFound      = errors.New("command Iptables not found")
	ErrIptablesNotMatch      = errors.New("No chain/target/match by that name")
	ErrRuleInfoChainNotFound = errors.New("Chain is required in the rule")
	ErrRuleInfoChainExist    = errors.New("rule to append has exist")
	ErrRuleInfoWrongAction   = errors.New("illegal action in rule")
	initOnce                 sync.Once
)

//util module
//to be imgrated
func sliceContainString(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func init() {
	initCheck()
}

func initDependencies() {
	detectIptables()
}

func initCheck() error {
	initOnce.Do(initDependencies)

	if IptablesPath == "" {
		return ErrIptablesNotFound
	}
	supportXlock = exec.Command(IptablesPath, "--wait", "-L", "-n").Run() == nil
	return nil
}

//Check if iptables command exists
func detectIptables() {
	path, err := exec.LookPath("iptables")
	if err != nil {
		fmt.Println(err)
	}
	IptablesPath = path
}

func GetVersion() (major, minor, micro int) {
	out, err := exec.Command(IptablesPath, "--version").CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	major, minor, micro = parseVersionNumber(string(out))
	return
}

func parseVersionNumber(input string) (major, minor, micro int) {
	re := regexp.MustCompile(`v\d*.\d*.\d*`)
	line := re.FindString(input)
	fmt.Sscanf(line, "v%d.%d.%d", &major, &minor, &micro)
	return
}

//--- RuleInfo ---
func (ri *RuleInfo) DefaultTable() {
	if ri.Table == nil {
		ri.Table = &FilterTable
	}
}

//--- Command ---
//execute iptables comamnd in raw args
func (ipt IPTable) raw(args ...string) (out []byte, err error) {
	out, err = exec.Command(IptablesPath, args...).CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return
}

//TODO current lock
// Table, Chain and Action are required in every command
func (ipt IPTable) Raw(ruleInfo *RuleInfo) ([]byte, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	ruleInfo.DefaultTable()
	args := ruleInfo.Args
	args = append([]string{"-t", string(ruleInfo.Table.Name), string(ruleInfo.Action), string(ruleInfo.Chain)}, args...)
	return ipt.raw(args...)
}

func (ipt IPTable) ValidateAndRun(ruleInfo *RuleInfo) error {
	validatableAction := []string{string(Append), string(Insert)}
	if sliceContainString(validatableAction, string(ruleInfo.Action)) {
		existsFlag, _ := ipt.RuleExists(ruleInfo)
		if existsFlag == true {
			return ErrRuleInfoChainExist
		}
	}
	out, err := ipt.Raw(ruleInfo)
	if err != nil {
		fmt.Println(err)
		return RuleError{
			Output: out,
		}
	}
	return nil
}

//--- CRUD ---
// 1. filtering
// 2. nat

func chainExistInTable(tableInfo *TableInfo, chain Chain) bool {
	for _, v := range tableInfo.Chains {
		if chain == v {
			return true
		}
	}
	return false
}

func targetExistInTable(tableInfo *TableInfo, target Target) bool {
	for _, v := range tableInfo.Targets {
		if target == v {
			return true
		}
	}
	return false
}

// iptables [-t table] -S [chain [rulenum]]
func (ipt IPTable) IptablesSave(tableInfo *TableInfo, chain Chain) (out []byte, err error) {
	if chain != "" {
		existFlag := chainExistInTable(tableInfo, chain)
		if existFlag != true {
			return nil, ErrIptablesNotMatch
		}
	}
	r := &RuleInfo{
		Table:  tableInfo,
		Chain:  chain,
		Action: Save,
	}
	return ipt.Raw(r)
}

//Q: support -C opt or not?
func (ipt IPTable) RuleExists(ruleInfo *RuleInfo) (bool, error) {
	//table and chain are required in ruleInfo
	if ruleInfo.Chain == "" {
		return false, ErrRuleInfoChainNotFound
	}
	ruleInfo.DefaultTable()

	ruleString := fmt.Sprintf("%s %s\n", ruleInfo.Chain, strings.Join(ruleInfo.Args, " "))
	existingRules, _ := ipt.IptablesSave(ruleInfo.Table, ruleInfo.Chain)
	return strings.Contains(string(existingRules), ruleString), nil
}

// default in filter table
//iptables [-t table] -P chain target
func (ipt IPTable) policySet(chain Chain, target Target) error {
	legalFlag := chainExistInTable(&FilterTable, chain) && targetExistInTable(&FilterTable, target)
	if legalFlag != true {
		return ErrIptablesNotMatch
	}

	r := &RuleInfo{
		Chain:  chain,
		Action: PolicySet,
		Args:   []string{string(target)},
	}
	out, err := ipt.Raw(r)
	if err != nil {
		fmt.Println(err)
		return RuleError{
			Output: out,
		}

	}
	return nil
}

func (ipt IPTable) ForwardPolicySet(target Target) error {
	return ipt.policySet(Forward, target)
}

//TODO App: Packet filtering
// match ip / tcp
// input/output? * accept/block? * tcp/udp? * ip/mask? * sport/dport? = ? rules

func parseIpAddress(s string) (string, error) {
	if strings.Contains(s, "/") == false {
		s = fmt.Sprintf("%s/32", s)
	}
	_, ipnet, err := net.ParseCIDR(s)
	if err != nil {
		return "", err
	}
	return ipnet.String(), nil
}

func IsStatefulAction(action Action) bool {
	for _, v := range statefulAction {
		if action == v {
			return true
		}
	}
	return false
}

func (ipt IPTable) InboundIp(action Action, source string, target Target) error {
	if IsStatefulAction(action) == false {
		return ErrRuleInfoWrongAction
	}
	if targetExistInTable(&FilterTable, target) == false {
		return ErrIptablesNotMatch
	}
	s, err := parseIpAddress(source)
	if err != nil {
		return err
	}
	r := &RuleInfo{
		Table:  &FilterTable,
		Chain:  Input,
		Action: action,
		Args:   []string{"-s", s, "-j", string(target)},
	}
	//	fmt.Printf("going to VandRun: %v\n", r)
	return ipt.ValidateAndRun(r)
}

//todo
func (ipt IPTable) OutboundIp(action Action, destination string, target Target) error {
	if IsStatefulAction(action) == false {
		return ErrRuleInfoWrongAction
	}
	if targetExistInTable(&FilterTable, target) == false {
		return ErrIptablesNotMatch
	}
	d, err := parseIpAddress(destination)
	if err != nil {
		return err
	}
	r := &RuleInfo{
		Table:  &FilterTable,
		Chain:  Output,
		Action: action,
		Args:   []string{"-d", string(d), "-j", string(Accept)},
	}
	return ipt.ValidateAndRun(r)
}

func (ipt IPTable) PortPermit(action Action, port int, target Target) error {
	if IsStatefulAction(action) == false {
		return ErrRuleInfoWrongAction
	}
	if targetExistInTable(&FilterTable, target) == false {
		return ErrIptablesNotMatch
	}
	r := &RuleInfo{
		Table:  &FilterTable,
		Chain:  Input,
		Action: action,
		Args:   []string{"-p", "tcp", "-m", "tcp", "--dport", strconv.Itoa(port), "-j", string(target)},
	}
	return ipt.ValidateAndRun(r)
}

func (ipt IPTable) Drop()    {}
func (ipt IPTable) Forward() {}

func (ipt IPTable) PreRouting()  {}
func (ipt IPTable) PostRouting() {}

//--- Application ---

//TODO  Accounting
//TODO  Connection tracking
//TODO  Packet mangling
//TODO  NAT network address translation
//TODO  Masquerading
//TODO  Port Forwarding
//TODO  Load balancing

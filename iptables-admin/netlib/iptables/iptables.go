package iptables

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"sync"
)

type Table string

type Chain string

type Policy string

type Action string

//iptables entry
type IPTable struct{}

type TableInfo struct {
	Name   Table
	Chains []Chain
}

// content of rule for iptables command
type RuleInfo struct {
	Table  *TableInfo
	chain  Chain
	Action Action
	Args   []string
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

	Accept Policy = "ACCEPT"
	Drop   Policy = "DROP"

	Append      Action = "-A"
	Delete      Action = "-D"
	Insert      Action = "-I"
	List        Action = "-L"
	NewChain    Action = "-N"
	DeleteChain Action = "-X"
	PolicySet   Action = "-P"
)

var (
	NatTable            TableInfo = TableInfo{Name: Nat, Chains: []Chain{Output, PreRouting, PreRouting}}
	FilterTable         TableInfo = TableInfo{Name: Filter, Chains: []Chain{Forward, Input, Output}}
	MangleTable         TableInfo = TableInfo{Name: Mangle, Chains: []Chain{Forward, Input, Output, PostRouting, PreRouting}}
	IptablesPath        string
	supportXlock        = false
	ErrIptablesNotFound = errors.New("command Iptables not found")
	initOnce            sync.Once
)

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

//--- Command ---
//execute iptables comamnd in raw args
func (ipt IPTable) raw(args ...string) {
	out, err := exec.Command(IptablesPath, args...).CombinedOutput()
	if err != nil {
		fmt.Println("in ipt.raw")
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

// Table is required in every command
// TODO switch args
func (ipt IPTable) Raw(ruleInfo RuleInfo) {
	table := ruleInfo.Table.Name
	if table == "" {
		table = Filter
	}

	args := ruleInfo.Args
	args = append([]string{"-t", string(table)}, args...)
	ipt.raw(args...)
}

//--- CRUD ---
// 1. filtering
// 2. nat

func (ipt IPTable) IptablesSave(TableInfo *TableInfo) {
	r := &RuleInfo{
		Table: TableInfo,
		Args:  []string{"-S"},
	}
	ipt.Raw(*r)
}

//--- Application ---
//TODO Packet filtering
//TODO  Accounting
//TODO  Connection tracking
//TODO  Packet mangling
//TODO  NAT network address translation
//TODO  Masquerading
//TODO  Port Forwarding
//TODO  Load balancing

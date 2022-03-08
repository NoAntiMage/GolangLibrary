package iptables

import (
	"fmt"
	"strconv"
	"testing"
)

var ()

func TestGetVersion(t *testing.T) {
	mj, mn, mc := parseVersionNumber("iptables v1.4.21")
	if mj != 1 || mn != 4 || mc != 21 {
		t.Fatal("Failed to parse version numbers")
	}
}

func TestIptableSave(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	ipt := IPTable{}
	tables := []TableInfo{FilterTable, NatTable, MangleTable}
	for _, table := range tables {
		chains := table.Chains
		for _, chain := range chains {
			_, err := ipt.IptablesSave(&table, chain)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

func TestPolicySet(t *testing.T) {
	ipt := IPTable{}
	chain := Forward
	// Drop
	policy := Drop
	r := &RuleInfo{
		//		Table:  &iptables.FilterTable,
		Chain:  chain,
		Action: PolicySet,
		Args:   []string{string(policy)},
	}

	ipt.ForwardPolicySet(policy)
	existsFlag, err := ipt.RuleExists(r)
	if err != nil {
		t.Fatal(err)
	}
	if existsFlag != true {
		t.Fatal("iptables policy setting does not work")
	}

	//Accept
	policy = Accept
	r.Args = []string{string(policy)}
	ipt.ForwardPolicySet(policy)
	existsFlag, err = ipt.RuleExists(r)
	if err != nil {
		t.Fatal(err)
	}
	if existsFlag != true {
		t.Fatal("iptables policy setting does not work")
	}
}

func TestInboundIp(t *testing.T) {
	ipt := IPTable{}
	source1 := "192.168.254.254"
	err := ipt.InboundIp(Append, source1, Drop)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		ipt.InboundIp(Delete, source1, Drop)
	}()

	r := &RuleInfo{
		Chain: Input,
		Args:  []string{"-s", source1 + "/32", "-j", string(Drop)},
	}
	existFlag, err := ipt.RuleExists(r)
	if existFlag == false || err != nil {
		t.Fatal("Inbound Ip does not work")
	}

	source2 := "192.168.253.0/24"
	err = ipt.InboundIp(Append, source2, Drop)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		ipt.InboundIp(Delete, source2, Drop)
	}()
	r = &RuleInfo{
		Chain: Input,
		Args:  []string{"-s", source2, "-j", string(Drop)},
	}
	existFlag, err = ipt.RuleExists(r)
	if existFlag == false || err != nil {
		t.Fatal("Inbound ipnet does not work")
	}

}

func TestOutboundIp(t *testing.T) {}

func TestPortPermit(t *testing.T) {
	ipt := IPTable{}
	targetPort := 30000
	ipt.PortPermit(Append, targetPort, Drop)
	defer func() {
		ipt.PortPermit(Delete, targetPort, Drop)
	}()

	r := &RuleInfo{
		Table:  &FilterTable,
		Chain:  Input,
		Action: Append,
		Args:   []string{"-p", "tcp", "-m", "tcp", "--dport", strconv.Itoa(targetPort), "-j", string(Drop)},
	}
	existFlag, err := ipt.RuleExists(r)
	if existFlag == false || err != nil {
		t.Fatal("PortPermit does not work")
	}
}

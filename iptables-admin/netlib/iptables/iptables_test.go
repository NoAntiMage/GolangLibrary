package iptables

import (
	"fmt"
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

// TODO
func TestInboundIp()
func TestOutboundIp()
func TestPortPermit()

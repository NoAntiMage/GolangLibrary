package main

import (
	"fmt"
	"netrouter/netlib/iptables"
)

func main() {
	ipt := &iptables.IPTable{}
	natTable := iptables.NatTable
	fmt.Println(natTable)
	ipt.IptablesSave(&natTable)
}

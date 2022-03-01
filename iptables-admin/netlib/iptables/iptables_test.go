package iptables

import (
	"testing"
)

func TestGetVersion(t *testing.T) {
	mj, mn, mc := parseVersionNumber("iptables v1.4.21")
	if mj != 1 || mn != 4 || mc != 21 {
		t.Fatal("Failed to parse version numbers")
	}
}

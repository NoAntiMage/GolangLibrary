package main

import (
	"fmt"
	"net"
)

func main() {
	IPAddrsGetInternal()
}

func IPAddrsGetInternal() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "0.0.0.0"
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				//return ipnet.IP.String()
			}
		}
	}
	return "0.0.0.0"
}

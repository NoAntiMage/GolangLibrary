package main

import (
	// "bufio"
	"fmt"
	"net"
)

// nslookup dig

func main() {
	// 查找DNS A记录
	iprecords, _ := net.LookupIP("baidu.com")
	for _, ip := range iprecords {
		fmt.Println("ip: ", ip)
	}

	// 查找 DNS CNAME记录
	cname, _ := net.LookupCNAME("www.baidu.com")
	fmt.Println("cname: ", cname)

	// 查找 DNS PTR记录
	ptr, _ := net.LookupAddr("8.8.8.8")
	for _, ptrval := range ptr {
		fmt.Println("ptr: ", ptrval)
	}
	//查找DNS NS记录
	namesever, _ := net.LookupNS("baidu.com")
	for _, ns := range namesever {
		fmt.Println("ns记录", ns)
	}
	//查找DNS MX记录
	mxrecords, _ := net.LookupMX("baidu.com")
	for _, mx := range mxrecords {
		fmt.Println("mx:", mx)
	}
	//查找DNS TXT记录
	txtrecords, _ := net.LookupTXT("baidu.com")
	for _, txt := range txtrecords {
		fmt.Println("txt:", txt)
	}
}

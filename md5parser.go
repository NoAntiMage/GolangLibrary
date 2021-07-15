package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

func md5Parse(bs []byte) string {
	m := md5.New()
	m.Write(bs)
	s := hex.EncodeToString(m.Sum(nil))
	result := strings.ToUpper(s)
	return result
}

func GetMD5Pwd(salt string, pwd string) string {
	d := []byte(salt + pwd)
	result := md5Parse(d)
	return result
}

func EncodeStr(str string) string {
	t := time.Now().Format("2006-01-02 15:04:05")
	d := []byte(str + t)
	result := md5Parse(d)
	return result
}

func main() {
	s := GetMD5Pwd("81ab5a8d-0a07-4152-9525-2ba36a9477f4", "password")
	fmt.Println(s)
	result := EncodeStr("string")
	fmt.Println(result)
}
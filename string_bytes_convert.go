package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func String2Bytes(s string) []byte {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func main() {
	s := "hello wujimaster"

	b1 := []byte(s)
	fmt.Println(b1)
	s1 := string(b1)
	fmt.Println(s1)

	b2 := String2Bytes(s1)
	fmt.Println(b2)
	s2 := Bytes2String(b2)
	fmt.Println(s2)
}

/*
reference:
src/runtime/string.go
func stringtoslicebyte()
func slicebytetostring()
*/

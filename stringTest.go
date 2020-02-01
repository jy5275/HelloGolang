package main

import (
	"fmt"
	"os"
)

// 自定义类型, 便于定义将方法包
// Bytesize 底层仍是 float64, 只是多了一包方法
type ByteSize float64

var (
	home   = os.Getenv("HOME")
	user   = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
)

func (b ByteSize) String() string {
	switch {
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	default:
		return fmt.Sprintf("%.2fB", b)
	}
}

func stringTest() {
	city := "美因河畔\x80法兰克福"
	cityr := []rune(city) // rune 就是 Golang 的 char
	fmt.Printf("%T, %q\n", cityr, cityr[:4])
	for pos, ch := range cityr {
		fmt.Printf("%#U starts at byte pos %v\n", ch, pos)
	}
	var t interface{} = cityr[2]
	r, ok := t.(int) // 安全的类型断言
	if ok {
		fmt.Println("r:", r)
	}
	switch t := t.(type) {
	case bool:
		fmt.Printf("boolean %t\n", t)
	case int:
		fmt.Printf("integer %d\n", t)
	case *bool:
		fmt.Printf("ptr to boolean %t\n", *t)
	case *int:
		fmt.Printf("ptr to integer %d\n", *t)
	case string:
		fmt.Printf("string: %s\n", t)
	case rune:
		fmt.Printf("rune(char) %#U\n", t)
	default:
		fmt.Println("Unknown type")
	}
	var b ByteSize = 12341
	fmt.Println(b)
	fmt.Println(home)
	fmt.Println(os.Args)
}

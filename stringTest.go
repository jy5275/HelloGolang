package main

import (
	"fmt"
	"os"
)

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
	city := "第聂伯彼得罗\x80夫斯克"
	cityr := []rune(city)
	fmt.Printf("%T, %q\n", cityr, cityr[:4])
	for pos, ch := range cityr {
		fmt.Printf("%#U starts at byte pos %v\n", ch, pos)
	}
	var t interface{} = city
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
	}
	var b ByteSize = 12341
	fmt.Println(b)
	fmt.Println(home)
	fmt.Println(os.Args)
}

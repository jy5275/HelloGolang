package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/jy5275/HelloGolang/morestrings"
	"io/ioutil"
	"net"
	"os"
)

func netWorkTest() {
	addr, e := net.ResolveTCPAddr("tcp4", "baidu.com:80")
	CheckErr(e)
	conn, e := net.DialTCP("tcp", nil, addr)
	CheckErr(e)
	_, e1 := conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	CheckErr(e1)
	res, e := ioutil.ReadAll(conn)
	CheckErr(e)
	fmt.Println(string(res))
}

func main() {
	//p.Init1()
	fmt.Println(morestrings.ReverseRunes("Reversed hello"))
	fmt.Println(cmp.Diff("Hello, World", "Hello, Go"))
	netWorkTest()
	funcObj()
}

// CheckErr ...
func CheckErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

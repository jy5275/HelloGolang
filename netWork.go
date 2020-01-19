package main

import (
	"fmt"
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

// CheckErr ...
func CheckErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

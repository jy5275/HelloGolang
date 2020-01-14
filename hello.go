package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/jy5275/hello/morestrings"
	_ "github.com/jy5275/hello/p"
)

func main() {
	//p.Init1()
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello, World", "Hello, Go"))
}

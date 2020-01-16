package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/jy5275/HelloGolang/morestrings"
	_ "github.com/jy5275/HelloGolang/p"
)

func main() {
	//p.Init1()
	fmt.Println(morestrings.ReverseRunes("Reversed hello"))
	fmt.Println(cmp.Diff("Hello, World", "Hello, Go"))
}

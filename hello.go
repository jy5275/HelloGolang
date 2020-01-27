/* Package main is a test package
 * Pretend there are some words...
 */
package main

import (
	"fmt"
	"./adv"
	"github.com/google/go-cmp/cmp"
	"github.com/jy5275/HelloGolang/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Reversed hello"))
	fmt.Println(cmp.Diff("Hello, World", "Hello, Go"))
	//netWorkTest()
	//funcObj()
	adv.fibTest()
	//stringTest()
	//NewMakeTest()
	//array2D(3, 5)
	//mapTest()
	//deferTest()
	//MarshalTest()
	//UnmarshalTest()
	//ifTest()
	//sTest()
}

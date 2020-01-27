/* Package main is a test package
 * Pretend there are some words...
 */
package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/jy5275/HelloGolang/adv"
	"github.com/jy5275/HelloGolang/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Reversed hello"))
	fmt.Println(cmp.Diff("Hello, World", "Hello, Go"))
	//netWorkTest()
	//funcObj()
	//adv.FibTest()
	adv.CalcMain()
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

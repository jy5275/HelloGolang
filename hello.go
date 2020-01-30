/* Package main is a test package
 * Pretend there are some words...
 */
package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/jy5275/HelloGolang/adv"
	"github.com/jy5275/HelloGolang/morestrings"
	_ "net/http/pprof"
	"runtime"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Reversed hello"))
	fmt.Println(cmp.Diff("Hello, World", "Hello, Go"))
	//adv.NetWorkTest()
	//funcObj()
	//MapTest()
	//stringTest()
	//NewMakeTest()
	//array2D(3, 5)
	//mapTest()
	//deferTest()
	//adv.MarshalTest()
	//adv.UnmarshalTest()
	//ifTest()
	//sTest()

	// Goroutine pratice
	//adv.FibTest()
	//adv.CalcMain()
	//adv.CtxMain()
	adv.SquareAndCube(589)
	//adv.PCMain()
	fmt.Println(runtime.NumGoroutine())
}

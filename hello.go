/* Package main is a test package
 * Pretend there are some words...
 */
package main

import (
	"fmt"
	_ "net/http/pprof"
	"runtime"
)

func main() {
	//adv.NetWorkTest()
	//funcObj()
	//stringTest()
	//NewMakeTest()
	deferTest()
	//adv.MarshalTest()
	//adv.UnmarshalTest()
	//ifTest()
	//sTest()

	// Goroutine pratice
	//adv.FibTest()
	//adv.CalcMain()
	//adv.CtxMain()
	//adv.SquareAndCube(589)
	//adv.PCMain()
	fmt.Println(runtime.NumGoroutine())
}

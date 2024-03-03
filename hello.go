/* Package main is a test package
 * Pretend there are some words...
 */
package main

import (
	"fmt"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

func GMP() {
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())
	//debug.SetMaxThreads(10)
	var wg sync.WaitGroup
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			a := 0
			for i := 0; i < 1e6; i++ {
				a += 1
			}
			wg.Done()
		}()
		time.Sleep(100 * time.Millisecond)
	}
	wg.Wait()
}

func main() {
	GMP()

}

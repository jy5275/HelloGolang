package adv

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"runtime/trace"
	"strconv"
	"sync"
)

func count(wg *sync.WaitGroup) {
	a := 0
	for j := 0; j < 1e6; j++ {
		a += 1
		if j == 1e6/2 {
			bytes, _ := ioutil.ReadFile(`gmp.go`)
			inc, _ := strconv.Atoi(string(bytes))
			a += inc
		}
	}
	wg.Done()
}

func GMP() {
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(2)
	//debug.SetMaxThreads(10)
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	var wg sync.WaitGroup
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go count(&wg)
		// time.Sleep(50 * time.Millisecond)
	}
	wg.Wait()
}

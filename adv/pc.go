package adv

import "sync"

import "time"

type data int

func makeData() data {
	return 1
}
func consumData(d data) {

}

func PCMain() {
	buf := make(chan data)
	exit := make(chan struct{})
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() { // Producer
			for {
				select {
				case <-exit:
					wg.Done()
					return
				default:
					buf <- makeData()
				}
			}
		}()
	}
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() { // Consumer
			for {
				select {
				case <-exit:
					wg.Done()
					return
				default:
					consumData(<-buf)
				}
			}
		}()
	}

	time.Sleep(2 * time.Second)
	close(exit)
	wg.Wait() // wg 下等了两个 goroutine, 它们都结束 main 再继续
}

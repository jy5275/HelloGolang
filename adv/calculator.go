package adv

import "fmt"

// 通过通信来共享数据, 而不是通过共享数据来通信!
type Cmd struct {
	A        int
	B        int
	Sum      int
	doneChan chan int // 计算完毕的信号,具体值不重要
}

type Calculator struct {
	cQueue chan *Cmd
}

func NewCalculator() *Calculator {
	cl := &Calculator{cQueue: make(chan *Cmd, 10)}
	go cl.calculator()
	return cl
}

func (cl *Calculator) Sum(a, b int) int {
	c := cl.sumAsync(a, b) // 为每个请求创建一个 Cmd 实例
	_ = <-c.doneChan       // 监听实例 doneChan, Adder 会放东西
	return c.Sum
}

func (cl *Calculator) sumAsync(a, b int) *Cmd {
	c := &Cmd{A: a, B: b, doneChan: make(chan int)}
	cl.cQueue <- c // 放入全局 chan
	return c
}

func (cl *Calculator) calculator() {
	for c := range cl.cQueue { // 取自全局 chan
		c.Sum = c.A + c.B
		c.doneChan <- 1
	}
}

func CalcMain() {
	cl := NewCalculator()
	fmt.Println(cl.Sum(2, 3))
}

package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func fibTest() {
	c := make(chan int, 5)
	quit := make(chan int)
	go fibonacci(c, quit)
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Main thread begin")
	// Abruptly output 5+ nums, then follows the producer's pace
	for i := 0; i < 12; {
		select {
		case num := <-c:
			fmt.Print(num, " ")
			i++
		default:
			fmt.Print(". ")
		}
		time.Sleep(50 * time.Millisecond)
	}
	quit <- 0
	time.Sleep(10)
}

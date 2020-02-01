package adv

import (
	"fmt"
	"time"
)

// Send Only && Receive Only Channel
func fibonacci(c chan<- int, quit <-chan bool) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case t := <-quit:
			fmt.Println("quit", t)
			return
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func FibTest() {
	c := make(chan int, 5)
	quit := make(chan bool)
	go fibonacci(c, quit) // fib can send on `c` and receive on `quit`
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
	close(quit)
	time.Sleep(350 * time.Millisecond)
}

func digits(n int, dch chan<- int) {
	for n != 0 {
		digit := n % 10
		dch <- digit
		n /= 10
	}
	close(dch)
	// Once closed, range operation on this channel will get a false ret val2
}

func calcCubes(num int, cubeop chan<- int) {
	sum := 0
	dch := make(chan int)
	go digits(num, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func calcSquares(num int, sqop chan<- int) {
	sum := 0
	dch := make(chan int)
	go digits(num, dch)
	for digit := range dch {
		sum += digit * digit
	}
	sqop <- sum
}

func SquareAndCube(n int) {
	sqch := make(chan int)
	cubech := make(chan int)
	go calcCubes(n, cubech)
	go calcSquares(n, sqch)
	sqres, cuberes := <-sqch, <-cubech
	fmt.Println(sqres + cuberes)
}

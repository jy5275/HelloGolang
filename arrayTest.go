package main

import (
	"fmt"
)

// Not visible!!! Modify on local Copy
func modifyArray(b [5]int) {
	for i := range b {
		b[i] = 200 + i
	}
}

// Visible to the caller, who shares the same referred array
func modifySlice(b []int, ret *[]int) {
	for i := range b {
		b[i] = 200 + i
		//*ret = append(*ret, b[i])	// Bad!
		// 不要在方法里给切片append，因为这会造成底层数组的改变
		// 不如返回一个新的slice
	}
}

func NewMakeTest() {
	var p *[5]int
	data := [...]int{0, 1, 2, 3, 4, 5, 6} // Array
	slice := data[1:4:5]
	_ = slice
	p = new([5]int) // Assign an array[5]
	modifyArray(*p)
	fmt.Println(p, len(*p), cap(*p))

	a := make([]int, 5)
	ret := make([]int, 0)
	modifySlice(a, &ret) // Not safe! modify on another `a`
	s := a[:2]
	fmt.Printf("%d, %d, %v\n", len(s), cap(s), s)
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Printf("%d, %d, %v\n", len(s), cap(s), s)
		fmt.Print(a, len(a), cap(a), "\n\n")
	}
}

func array2D(YSize, XSize int) {
	picture := make([][]uint8, YSize)
	pixels := make([]uint8, XSize*YSize)
	for i := range picture {
		picture[i], pixels = pixels[:XSize], pixels[XSize:]
	}
}

func PrintArray(e ...int) {
	for _, i := range e {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

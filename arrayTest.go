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
		*ret = append(*ret, b[i])
	}
}

func MyAppend(in []int, data int) []int {
	//slice := p
	out := append(in, data)
	//p = newslice
	return out
}

func NewMakeTest() {
	var p *[5]int
	p = new([5]int) // Assign an array[5]
	modifyArray(*p)
	fmt.Println(p, len(*p), cap(*p))

	v := make([]int, 5)
	ret := make([]int, 0)
	modifySlice(v, &ret)
	exvs := v[:2]
	fmt.Printf("%T, %d, %d, %v\n", exvs, len(exvs),
		cap(exvs), exvs)
	for i := 0; i < 10; i++ {
		exvs = MyAppend(exvs, i)
		//exvs = append(exvs, i)
		fmt.Printf("%d, %d, %v\n", len(exvs),
			cap(exvs), exvs)
		fmt.Print(v, len(v), cap(v), "\n\n")
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

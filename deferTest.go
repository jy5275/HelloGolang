package main

import "fmt"

func f1() (res int) {
	defer func() {
		res++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t += 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r += 5
	}(r)
	return 1
}

func deferTest() {
	fmt.Println(f1(), f2(), f3())
}

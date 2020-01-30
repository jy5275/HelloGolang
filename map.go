package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

type Stu struct {
	id   *int
	name string
}

func MapTest() {
	m := make(map[Stu]bool) // m 实质上是个引用！不 make 就 nil
	for {
		var stu Stu
		tmpint := 9
		fmt.Scanf("%s", &stu.name)
		stu.id = &tmpint
		if stu.name == "ST" {
			break
		}
		m[stu] = true
	}
	fmt.Println(m)
}

// WordCount ...
func WordCount(s string) map[string]int {
	strs := strings.Fields(s)
	m := make(map[string]int)
	for i := range strs {
		m[strs[i]]++
	}
	newm := make(map[string]int)
	for k, v := range m {
		newm[k] = v
	}
	return m
}

func mapTest() {
	wc.Test(WordCount)
}

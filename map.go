package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

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

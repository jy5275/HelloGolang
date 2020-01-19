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
	return m
}

func mapTest() {
	wc.Test(WordCount)
}

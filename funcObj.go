package main

import (
	"fmt"
)

type Adder interface {
	add(string) int
}
type handler func(name string) int

// 一旦通过函数名调用方法, 函数名(函数指针)就变成了函数对象
func (h handler) add(name string) int {
	return h(name) + 10
}

// process函数需要Adder接口类型参数
func process(a Adder) {
	fmt.Println("process:", a.add("taozs"))
}

type myint int // 基本类型别名 myint 也实现了 Adder 接口
func (i myint) add(name string) int {
	return len(name) + int(i)
}

func funcObj() {
	var my handler = func(name string) int {
		return len(name)
	}
	var doubler handler = func(name string) int {
		return len(name) * 2
	}
	fmt.Println(my("taozs"))                   // 调用函数 5, my 是函数指针
	fmt.Println(my.add("taozs"))               // 调用函数对象的方法 15, my是函数对象
	fmt.Println(handler(doubler).add("taozs")) // doubler 显式转成 handler 函数对象后调用对象的add方法20
	// 以下是针对接口Adder的调用
	process(my)               // 5
	process(handler(doubler)) // process 接受参数类型是 handler，强转 20
	process(myint(8))         //实现Adder接口不仅可以是函数也可以是结构体 13

}

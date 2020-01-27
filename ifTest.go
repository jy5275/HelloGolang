package main

import "fmt"

type Student interface {
	study()
	exam()
}
type Employee interface {
	work()
}
type PhD struct {
	Student
	Employee
}
type Coder struct {
	Employee
}

func (PhD) work() {
	fmt.Println("Banzhuaning...")
}
func (PhD) study() {
	fmt.Println("PhD is studying...")
}
func (PhD) exam() {
	fmt.Println("PhD taking exam...")
}
func (Coder) work() {
	fmt.Println("Coding...")
}

func ifTest() {
	//var s Student
	p := PhD{}
	s := Student(p)
	//p := new(PhD)
	if p, ok := s.(PhD); ok {
		fmt.Println(p)
	} else {
		fmt.Println("Not a PhD")
	}
	//.play()
}

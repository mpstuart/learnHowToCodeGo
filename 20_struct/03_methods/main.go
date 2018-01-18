package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() (string, int) {
	return p.first + p.last, p.age
}

func main() {
	p1 := person{"James ", "Bond ", 20}
	p2 := person{"Miss ", "Moneypenny ", 18}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}

package main

import "fmt"

func main() {
	var x int
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
anonymous function
a function without a name

func expression
assigning a func to a variable
*/

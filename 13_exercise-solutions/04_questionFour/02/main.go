package main

import "fmt"

func main() {
	var numOne int
	var numTwo int
	fmt.Println("Enter a number: ")
	fmt.Scan(&numOne)
	fmt.Println("Enter a Larger number: ")
	fmt.Scan(&numTwo)
	fmt.Println(numTwo, " / ", numOne, " = ", numTwo/numOne)
	fmt.Println("Remainder = ", numTwo%numOne)
}

package main

import "fmt"

func main() {

	var x string
	var y string
	fmt.Printf("Enter a small number> ")
	fmt.Scan(&x)
	fmt.Printf("Enter a larger number> ")
	fmt.Scan(&y)
	fmt.Println("This is the smaller number", x)
	fmt.Println("This is the larger number", y)

}

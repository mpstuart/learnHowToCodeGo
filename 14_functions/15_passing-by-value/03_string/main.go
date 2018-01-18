package main

import "fmt"

func main() {

	name := "Michael"
	fmt.Println(name) // Michael

	changeMe(name)

	fmt.Println(name)
}

func changeMe(z string) {
	fmt.Println(z) // Michael
	z = "Rocky"
	fmt.Println(z) // Rocky
}

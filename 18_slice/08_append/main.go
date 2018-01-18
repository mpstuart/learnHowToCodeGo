package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	//3 is the length
	//5 is capacity

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting = append(greeting, "suprabadham")

	fmt.Println(greeting[3])
}

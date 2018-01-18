package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"zero":  "Good morning!",
		"one":   "Bonjour!",
		"two":   "Buenos dias!",
		"three": "Bongiorno!",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, "two")
	fmt.Println(myGreeting)

	myGreetin := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(myGreetin)
	delete(myGreetin, 2)
	fmt.Println(myGreetin)
}

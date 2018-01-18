package main

import "fmt"

func main() {

	counter := 0
	for x := 0; x < 1000; x++ {
		if x%3 == 0 {
			counter += x
		} else if x%5 == 0 {
			counter += x
		}
	}

	fmt.Println(counter)
}

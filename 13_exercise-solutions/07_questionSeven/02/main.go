package main

import "fmt"

func main() {
	var counter int
	for x := 0; x < 1000; x++ {
		if x%5 == 0 {
			counter += x
		} else if x%3 == 0 {
			counter += x
		}
	}
	fmt.Println(counter)
}

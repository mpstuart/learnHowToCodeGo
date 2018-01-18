package main

import "fmt"

func main() {
	for x := 0; x <= 100; x++ {
		if x%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if x%5 == 0 {
			fmt.Println("Buzz")
		} else if x%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(x)
		}
	}
}

package main

import "fmt"

func incro(n ...int) int {
	var total int
	for _, v := range n {
		total += v
	}
	return total
}
func main() {
	var x int
	for i := 0; i < 1000; i++ {
		if i%5 == 0 {
			x += incro(i)
		} else if i%3 == 0 {
			x += incro(i)
		}
	}
	fmt.Println(x)
}

/*
add to slice or map then ... total of map
add different packages for the different functions
*/

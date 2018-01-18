package main

import "fmt"

func half(i int) (int, bool) {
	return i / 2, i%2 == 0
}

func main() {
	h, e := half(4)
	fmt.Println(h, e)
	//same top is substitution
	fmt.Println(half(5))
}

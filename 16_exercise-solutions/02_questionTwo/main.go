package main

import "fmt"

func main() {

	half := func(i int) (int, bool) {
		return i / 2, i%2 == 0
	}
	h, e := half(5)
	fmt.Println(h, e)
	//same the top is substitution
	fmt.Println(half(5))
}

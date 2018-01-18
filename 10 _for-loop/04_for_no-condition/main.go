package main

import "fmt"

func main() {
	i := 0
	for { // no condition will run forever
		fmt.Println(i)
		i++
	}
}

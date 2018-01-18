package main

import "fmt"

var x rune = 'a'
var y int32 = 'b'

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(string(x))
	fmt.Println(string(y))
	// conversion: rune to string

}

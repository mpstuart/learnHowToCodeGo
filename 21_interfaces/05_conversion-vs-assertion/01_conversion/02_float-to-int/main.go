package main

import "fmt"

var x int = 12
var y float64 = 12.1230123

func main() {
	fmt.Println(int(y) + x)
	// conversion: float64 to int
}

package main

import "fmt"

var x int = 12
var y float64 = 12.1230123

func main() {
	fmt.Println(y + float64(x))
	// conversion: int to float64
}

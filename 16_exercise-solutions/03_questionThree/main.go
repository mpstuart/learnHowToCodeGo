package main

import "fmt"

func max(nums ...int) int {
	var largest int
	for _, v := range nums {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	fmt.Println(max(
		4,
		7,
		9,
		123,
		543,
		23,
		435,
		53,
		125,
	))
}

package main

import "fmt"

func max_value(x int, y int) int {
	var max int
	if x > y {
		max = x
	} else {
		max = y
	}
	return max
}

func main() {
	var x1 int = 5
	var y1 int = 2
	result := max_value(x1,y1)
	fmt.Println(result)
}
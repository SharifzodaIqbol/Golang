package main

import "fmt"

func more(x int, y int) bool {
	if x > y {
		return true
	} else {
		return false
	}
}

func main() {
	var x1 int = 5
	var y1 int = 2
	result := more(x1,y1)
	fmt.Println(result)
}
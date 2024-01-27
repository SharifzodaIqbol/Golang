package main

import (
	"fmt"	
)

func maxAreaCheck(length int, width int, limit int) bool {
	if area := length * width; area < limit {
		return true
	} else{
		return false
	}
}

func main() {

	length := 7
	width := 5
	limit := 34
	result := maxAreaCheck(length,width,limit)
	fmt.Println(result) 
}
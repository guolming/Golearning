package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 3, 4, 5, 6, 7, 3, 5, 6}
	result := -1
	for _, item := range arr {
		if result < 0 {
			result = item
		} else {
			result = result ^ item
		}
	}
	fmt.Println(result)
}

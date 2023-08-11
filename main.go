package main

import (
	"fmt"
)

func Pyramid(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		subArray := make([]int, i+1)
		for j := 0; j <= i; j++ {
			subArray[j] = 1
		}
		result[i] = subArray
	}
	return result
}

func main() {
	fmt.Println(Pyramid(0))
	fmt.Println(Pyramid(1))
	fmt.Println(Pyramid(2))
	fmt.Println(Pyramid(3))
}

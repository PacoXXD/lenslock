package main

// you can also use imports, for example:

import (
	"fmt"
	"sort"
)

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X []int, Y []int) int {
	// Implement your solution here
	sort.Ints(X)
	max := 0
	for i := 1; i < len(X); i++ {
		dist := X[i] - X[i-1]
		if dist > max {
			max = dist
		}
	}
	return max
}

func main() {
	X := []int{1, 8, 7, 3, 4, 1, 8}
	Y := []int{6, 4, 1, 8, 5, 1, 7}
	fmt.Println(Solution(X, Y))
}

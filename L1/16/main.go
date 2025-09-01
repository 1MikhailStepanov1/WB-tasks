package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("%v", quickSort([]int{55, 235, 66, 9, 10, 2, 5}))
}

func quickSort(input []int) []int {
	if len(input) < 2 {
		return input
	}

	left := 0
	right := len(input) - 1
	pivotIdx := rand.Int() % len(input)
	input[pivotIdx], input[right] = input[right], input[pivotIdx]

	for i := range input {
		if input[i] < input[right] {
			input[i], input[left] = input[left], input[i]
			left++
		}
	}
	input[left], input[right] = input[right], input[left]

	quickSort(input[:left])
	quickSort(input[left+1:])
	return input
}

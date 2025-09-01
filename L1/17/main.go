package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(binarySearch(7, values, 0, len(values)-1))
}

func binarySearch(el int, input []int, lowIdx int, highIdx int) int {
	if lowIdx > highIdx {
		return -1
	}
	halfIdx := lowIdx + (highIdx-lowIdx)/2
	if input[halfIdx] == el {
		return halfIdx
	} else if el < input[halfIdx] {
		return binarySearch(el, input, lowIdx, halfIdx-1)
	} else {
		return binarySearch(el, input, halfIdx+1, highIdx)
	}
}

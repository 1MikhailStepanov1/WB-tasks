package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var values []int

	for i := 0; i < 10; i++ {
		values = append(values, rand.Intn(100))
	}
	fmt.Println("Start slice: ", values)
	elementToRemove := rand.Intn(10)
	for {
		if elementToRemove == 0 || elementToRemove == len(values)-1 {
			elementToRemove = rand.Intn(10)
		} else {
			break
		}
	}
	fmt.Println("Removed element index: ", elementToRemove)

	copy(values[elementToRemove:], values[elementToRemove+1:])
	values = values[:len(values)-1]
	fmt.Println("Finish slice: ", values)
}

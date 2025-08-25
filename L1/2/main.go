package main

import (
	"fmt"
	"sync"
)

func main() {
	array := [5]int{2, 4, 6, 8, 10}

	wg := new(sync.WaitGroup)
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(x int) {
			defer wg.Done()
			fmt.Println(x * x)
		}(array[i])
	}
	wg.Wait()
}

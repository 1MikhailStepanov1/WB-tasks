package main

import "fmt"

func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for data := range in {
			out <- data * 2
		}
	}()
	return out
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	in := make(chan int)
	go func() {
		defer close(in)
		for i := 0; i < len(data); i++ {
			in <- data[i]
		}
	}()
	for doubleData := range double(in) {
		fmt.Println(doubleData)
	}
}

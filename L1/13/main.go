package main

import "fmt"

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		return
	}

	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
}

/*
11 56
56 + 11 = 67
67 56
67 - 56 = 11
67 11
67 - 11 = 56
56 11
*/

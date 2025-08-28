package main

import "fmt"

func main() {
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	result := map[int][]float64{}

	for _, d := range data {
		result[int(d)/10*10] = append(result[int(d)/10*10], d)
	}
	fmt.Printf("%v", result)
}

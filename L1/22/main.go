package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, b := big.NewInt(10000001), big.NewInt(10000002)
	var sub, add, mul, div big.Int
	sub.Sub(a, b)
	add.Add(a, b)
	mul.Mul(a, b)
	div.Div(a, b)
	fmt.Println("Substraction: ", sub)
	fmt.Println("Addition: ", add)
	fmt.Println("Multiplication: ", mul)
	fmt.Println("Division: ", div)
}

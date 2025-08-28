package main

import (
	"flag"
	"fmt"
	"log"
	"math/bits"
)

func main() {
	number := flag.Uint("number", 5, "number for program")
	bitNumber := flag.Uint("bit", 0, "bit number to change")
	flag.Parse()

	if bits.Len(*number)-1 < int(*bitNumber) {
		log.Fatal("Number of bits in smaller than presented value")
	}

	fmt.Printf("Entered number: %b\n", *number)
	mask := uint(1) << *bitNumber
	fmt.Printf("Mask: %b\n", mask)
	result := *number ^ mask
	fmt.Printf("Result: %b\n", result)
}

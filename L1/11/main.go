package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	set1 := map[int]struct{}{}
	set2 := map[int]struct{}{}
	result := map[int]struct{}{}

	for i := 0; i < 2; i++ {
		scanner.Scan()

		parts := strings.Split(scanner.Text(), ",")
		for _, part := range parts {
			if num, err := strconv.Atoi(strings.TrimSpace(part)); err == nil {
				if i == 0 {
					set1[num] = struct{}{}
				} else {
					set2[num] = struct{}{}
				}
			}
		}
	}

	for el := range set1 {
		if _, exists := set2[el]; exists {
			result[el] = set1[el]
		}
	}

	for el := range result {
		fmt.Printf("%d\t", el)
	}
}

/*
Считываем первый слайс,
*/

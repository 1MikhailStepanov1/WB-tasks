package main

import (
	"fmt"
	"unicode"
)

func main() {
	symbols := map[rune]struct{}{}
	var text string
	_, err := fmt.Scan(&text)
	if err != nil {
		return
	}

	endFlag := true
	for _, s := range []rune(text) {
		s = unicode.ToLower(s)
		if _, exists := symbols[s]; !exists {
			symbols[s] = struct{}{}
		} else {
			endFlag = false
			break
		}
	}
	fmt.Println(endFlag)
}

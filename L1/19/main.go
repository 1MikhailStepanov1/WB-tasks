package main

import "fmt"

func main() {
	var text string
	_, err := fmt.Scan(&text)
	if err != nil {
		return
	}
	runeArray := []rune(text)
	for i := 0; i < len(runeArray); i++ {
		fmt.Printf("%s", string(runeArray[len(runeArray)-1-i]))
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		wordsArray := strings.Split(scanner.Text(), " ")
		for i := 0; i < len(wordsArray); i++ {
			fmt.Printf("%s ", wordsArray[len(wordsArray)-1-i])
		}
	}
}

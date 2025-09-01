package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	result := map[string]struct{}{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		result[strings.Trim(line, " ")] = struct{}{}
	}

	for line := range result {
		fmt.Printf("%s\t", line)
	}

}

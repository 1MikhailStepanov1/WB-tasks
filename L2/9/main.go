package main

import (
	"errors"
	"fmt"
	"os"
	"unicode"
)

func unpack(input string) (string, error) {
	inputRunes := []rune(input)
	result := make([]rune, 0)
	isShielded := false

	for i := 0; i < len(inputRunes); i++ {
		if unicode.IsLetter(inputRunes[i]) {
			result = append(result, inputRunes[i])
		} else if unicode.IsDigit(inputRunes[i]) && len(result) > 0 {
			if inputRunes[i] == '0' {
				return "", errors.New("can't use 0 number")
			}
			if isShielded {
				isShielded = false
				result = append(result, inputRunes[i])
			} else {
				lastLetter := result[len(result)-1]
				for j := 0; j < int(inputRunes[i]-'0')-1; j++ {
					result = append(result, lastLetter)
				}
			}
		} else if inputRunes[i] == '\\' {
			if isShielded {
				return "", errors.New("can't use \\ character")
			} else {
				isShielded = true
			}
		} else {
			return "", errors.New("invalid character sequence")
		}
	}

	if len(result) == 0 && len(input) > 0 {
		return "", errors.New("string doesn't contains any symbols to unpack")
	} else if isShielded {
		return "", errors.New("unclosed shielded symbol")
	}
	return string(result), nil
}

func main() {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		return
	}
	result, err := unpack(input)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(result)
}

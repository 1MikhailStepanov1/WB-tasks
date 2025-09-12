package main

import (
	"errors"
	"fmt"
	"os"
	"unicode"
)

var (
	ErrShieldOnlyNumbers  = errors.New("only numbers can be shielded")
	ErrDigitWithoutLetter = errors.New("digit cannot be first character without letter")
	ErrZeroNumber         = errors.New("can't use 0 number")
	ErrInvalidSequence    = errors.New("invalid character sequence")
	ErrUnclosedShield     = errors.New("unclosed shielded symbol")
	ErrNoSymbolsToUnpack  = errors.New("string doesn't contain any symbols to unpack")
)

func unpack(input string) (string, error) { //nolint:cyclop
	if input == "" {
		return "", nil
	}

	inputRunes := []rune(input)
	result := make([]rune, 0)
	isShielded := false

	for i := range inputRunes {
		current := inputRunes[i]

		switch {
		case isShielded:
			if current != '\\' {
				result = append(result, current)
				isShielded = false
			} else {
				return "", ErrShieldOnlyNumbers
			}
		case unicode.IsLetter(current):
			result = append(result, current)

		case current == '\\':
			isShielded = true

		case unicode.IsDigit(current):
			if len(result) == 0 {
				return "", ErrDigitWithoutLetter
			}
			if current == '0' {
				return "", ErrZeroNumber
			}

			lastLetter := result[len(result)-1]
			count := int(current-'0') - 1
			for range count {
				result = append(result, lastLetter)
			}

		default:
			return "", ErrInvalidSequence
		}
	}

	if isShielded {
		return "", ErrUnclosedShield
	}

	if len(result) == 0 {
		return "", ErrNoSymbolsToUnpack
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
		fmt.Println("Error:", err) //nolint:forbidigo
		os.Exit(1)
	}
	fmt.Println(result) //nolint:forbidigo
}

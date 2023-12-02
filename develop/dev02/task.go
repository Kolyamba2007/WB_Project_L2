package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var ErrInvStr = errors.New("некорректная строка")

func main() {
	str, err := unpack("a4bc2d5e")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Uncorrect string: %v\n", err)
		os.Exit(123)
	}

	fmt.Println(str)
}

func unpack(str string) (finalStr string, err error) {
	if str == "" {
		return
	}

	var lastRune rune

	for i, currentRune := range str {
		if unicode.IsDigit(currentRune) {
			if i == 0 || unicode.IsDigit(lastRune) {
				return "", ErrInvStr
			}

			finalStr += strings.Repeat(string(lastRune), int(currentRune-'0'))

			lastRune = currentRune

			continue
		}

		if !unicode.IsDigit(lastRune) && lastRune != 0 {
			finalStr += string(lastRune)
		}

		lastRune = currentRune
	}

	if !unicode.IsDigit(lastRune) {
		finalStr += string(lastRune)
	}

	return
}

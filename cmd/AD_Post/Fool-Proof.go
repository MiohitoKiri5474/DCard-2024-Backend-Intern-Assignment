package main

import (
	"strconv"
	"unicode"
)

func AgeCheck(str string) bool {
	for _, char := range str {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	num, _ := strconv.Atoi(str)
	return 1 <= num && num <= 100
}

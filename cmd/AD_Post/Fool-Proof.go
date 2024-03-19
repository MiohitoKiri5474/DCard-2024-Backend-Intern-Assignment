package main

import (
	"strconv"
	"unicode"
)

func AgeCheck(str string) error {
	for _, char := range str {
		if !unicode.IsDigit(char) {
			panic("non-digit input")
		}
	}
	if num, _ := strconv.Atoi(str); num < 1 || 100 < num {
		panic("out of range")
	}
	return nil
}

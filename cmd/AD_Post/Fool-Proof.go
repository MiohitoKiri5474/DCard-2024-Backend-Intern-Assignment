package main

import (
	"strconv"
	"unicode"

	"github.com/mikekonan/go-countries"
)

func CheckAge(str string) error {
	// check age is valid
	for _, char := range str {
		if !unicode.IsDigit(char) {
			panic("(age) non-digit input")
		}
	}
	if num, _ := strconv.Atoi(str); num < 1 || 100 < num {
		panic("(age) out of range")
	}
	return nil
}

func CheckGender(str string) error {
	// check gender is valid
	if str != "F" && str != "M" {
		panic("(gender) invalid gender")
	}
	return nil
}

func CheckCountry(str string) error {
	// check country code is valid
	tmp := country.Alpha2Code(str)
	if _, ok := country.ByAlpha2Code(tmp); !ok {
		panic("(country code) invalid ISO 3166-1 code")
	}
	return nil
}

func CheckPlatform(str string) error {
	// check platform is valid
	if str != "web" && str != "iso" && str != "android" {
		panic("(platform) invalid platform")
	}
	return nil
}

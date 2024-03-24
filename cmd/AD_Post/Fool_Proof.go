package main

import (
	"AD_Post/models"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/mikekonan/go-countries"
)

func CheckAge(str string) error {
	// check age is valid
	if str == "" {
		return nil
	}
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
	if str == "" {
		return nil
	}
	if str != "F" && str != "M" {
		panic("(gender) invalid gender")
	}
	return nil
}

func CheckCountry(str string) error {
	// check country code is valid
	if str == "" {
		return nil
	}
	tmp := country.Alpha2Code(str)
	if _, ok := country.ByAlpha2Code(tmp); !ok {
		panic("(country code) invalid ISO 3166-1 code")
	}
	return nil
}

func CheckPlatform(str string) error {
	// check platform is valid
	if str == "" {
		return nil
	}
	LowercaseStr := strings.ToLower(str)
	if LowercaseStr != "web" && LowercaseStr != "ios" && LowercaseStr != "android" {
		panic("(platform) invalid platform")
	}
	return nil
}

func CheckAgeInt(num int, name string) error {
	if num < 1 || 100 < num {
		errMsg := fmt.Sprintf("(%s) out of range", name)
		panic(errMsg)
	}
	return nil
}

func CheckJSon(input models.JsonParse) error {
	// Check Json conditions

	// Check age
	if err := CheckAgeInt(input.Conditions.AgeStart, "ageStart"); err != nil {
		panic(err.Error())
	}
	if err := CheckAgeInt(input.Conditions.AgeEnd, "ageEnd"); err != nil {
		panic(err.Error())
	}
	if input.Conditions.AgeEnd < input.Conditions.AgeStart {
		panic("(ageStart, ageEnd) range error")
	}

	// Check gender
	for _, gender := range input.Conditions.Gender {
		if err := CheckGender(gender); err != nil {
			panic(err.Error())
		}
	}

	// Check country
	for _, country := range input.Conditions.Country {
		if err := CheckCountry(country); err != nil {
			panic(err.Error())
		}
	}

	// Check platform
	for _, platform := range input.Conditions.Platform {
		if err := CheckPlatform(platform); err != nil {
			panic(err.Error())
		}
	}
	return nil
}

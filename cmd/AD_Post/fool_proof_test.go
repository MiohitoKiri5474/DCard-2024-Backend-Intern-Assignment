package main

import (
	"testing"
)

func TestCheckAge(t *testing.T) {
	// Test CheckAge, suppose panic will not occur
	if err := CheckAge("10"); err != nil {
		t.Error()
	}
}

func TestCheckAgeFailed(t *testing.T) {
	// Test CheckAge, suppose panic will occur
	defer func() {
		if r := recover(); r != nil {
			// Panic occurred as expected, test passed
			return
		}
		// If recover didn't catch a panic, test failed
		t.Error("Expected panic did not occur")
	}()
	_ = CheckAge("1000")
}

func TestCheckGender(t *testing.T) {
	// Test CheckGender, suppose panic will not occur
	if err := CheckGender("M"); err != nil {
		t.Error()
	}

	if err := CheckGender("F"); err != nil {
		t.Error()
	}
}

func TestCheckGenderFailed(t *testing.T) {
	// Test CheckGender, suppose panic will occur
	defer func() {
		if r := recover(); r != nil {
			// Panic occurred as expected, test passed
			return
		}
		// If recover didn't catch a panic, test failed
		t.Error("Expected panic did not occur")
	}()
	_ = CheckGender("1000")
}

func TestCheckCountry(t *testing.T) {
	// Test CheckCountry, suppose panic will not occur
	if err := CheckCountry("JP"); err != nil {
		t.Error()
	}

	if err := CheckCountry("US"); err != nil {
		t.Error()
	}

	if err := CheckCountry("TW"); err != nil {
		t.Error()
	}
}

func TestCheckCountryFailed(t *testing.T) {
	// Test CheckCountrt, suppose panic will occur
	defer func() {
		if r := recover(); r != nil {
			// Panic occurred as expected, test passed
			return
		}
		// If recover didn't catch a panic, test failed
		t.Error("Expected panic did not occur")
	}()
	_ = CheckCountry("1000")
}

func TestCheckPlatform(t *testing.T) {
	// Test CheckPlatform, suppose panic will not occur
	if err := CheckPlatform("ios"); err != nil {
		t.Error()
	}

	if err := CheckPlatform("web"); err != nil {
		t.Error()
	}

	if err := CheckPlatform("android"); err != nil {
		t.Error()
	}
}

func TestCheckPlatformFailed(t *testing.T) {
	// Test CheckPlatform, suppose panic will occur
	defer func() {
		if r := recover(); r != nil {
			// Panic occurred as expected, test passed
			return
		}
		// If recover didn't catch a panic, test failed
		t.Error("Expected panic did not occur")
	}()
	_ = CheckPlatform("1000")
}

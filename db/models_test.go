package models

import (
	"testing"
)

func TestCreate(t *testing.T) {
	BuildDB("123.db")
}

func TestCompressJSON(t *testing.T) {
	OriList := []string{"123", "456", "789"}
	ExpectRes := "123 456 789"
	if ExpectRes != CompressJSON(OriList) {
		t.Error("Value Error")
	}
}

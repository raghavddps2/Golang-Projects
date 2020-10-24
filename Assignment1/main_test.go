package main

import (
	"testing"
)

func TestGetOddOrEven(t *testing.T) {
	ans := getOddOrEven(10)
	if ans != "Even" {
		t.Errorf("Expected Even got %v ",ans)
	}
}

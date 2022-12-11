package main

import (
	_ "embed"
	"testing"
)

//go:embed test.txt
var test []byte

func TestPart1(t *testing.T) {
	const expected = 21
	ans := runPart1(test)
	if ans != expected {
		t.Errorf("Received: %d; Expected %d", ans, expected)
	}
}

func TestPart2(t *testing.T) {
	const expected = 8
	ans := runPart2(test)
	if ans != expected {
		t.Errorf("Received: %d; Expected %d", ans, expected)
	}
}

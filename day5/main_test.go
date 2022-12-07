package main

import (
	_ "embed"
	"testing"
)

//go:embed test.txt
var test []byte

func TestPart1(t *testing.T) {
	const expected = "CMZ"
	ans := runPart1(test)
	if ans != expected {
		t.Errorf("Received: %s; Expected %s", ans, expected)
	}
}

func TestPart2(t *testing.T) {
	const expected = "MCD"
	ans := runPart2(test)
	if ans != expected {
		t.Errorf("Received: %s; Expected %s", ans, expected)
	}
}

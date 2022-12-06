package main

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed test.txt
var test []byte

func TestPart1(t *testing.T) {
	const expected = 157
	fmt.Println(test)
	ans := runPart1(test)
	if ans != expected {
		t.Errorf("Received: %d; Expected %d", ans, expected)
	}
}

func TestPart2(t *testing.T) {
	const expected = 70
	fmt.Println(test)
	ans := runPart2(test)
	if ans != expected {
		t.Errorf("Received: %d; Expected %d", ans, expected)
	}
}

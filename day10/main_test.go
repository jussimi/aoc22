package main

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed test.txt
var test []byte

func TestPart1(t *testing.T) {
	const expected = 13140
	ans := runPart1(test)
	if ans != expected {
		t.Errorf("Received: %d; Expected %d", ans, expected)
	}
}

func TestPart2(t *testing.T) {
	lines := []string{
		"##..##..##..##..##..##..##..##..##..##..",
		"###...###...###...###...###...###...###.",
		"####....####....####....####....####....",
		"#####.....#####.....#####.....#####.....",
		"######......######......######......####",
		"#######.......#######.......#######.....",
	}

	expected := strings.Join(lines, "\n")

	ans := runPart2(test)
	if ans != expected {
		t.Errorf("\nReceived: \n%s\nExpected \n%s", ans, expected)
	}
}

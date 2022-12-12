package main

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"
)

//go:embed test.txt
var test []byte

func TestPart1(t *testing.T) {
	const expected = 13
	input := strings.Split(string(test), "\n\n")[0]
	ans := runPart1([]byte(input))
	if ans != expected {
		t.Errorf("Received: %d; Expected %d", ans, expected)
	}
}

func TestPart2(t *testing.T) {
	var expected = []int{1, 36}

	for i, test := range strings.Split(string(test), "\n\n") {
		testname := fmt.Sprintf("Test %d", i)
		t.Run(testname, func(t *testing.T) {
			ans := runPart2([]byte(test))
			if ans != expected[i] {
				t.Errorf("Received: %d; Expected %d", ans, expected[i])
			}
		})
	}
}

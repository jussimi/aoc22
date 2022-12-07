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
	var expected = []int{7, 5, 6, 10, 11}

	for i, line := range strings.Split(string(test), "\n") {
		testname := fmt.Sprintf("Line %d", i)
		t.Run(testname, func(t *testing.T) {
			ans := runPart1([]byte(line))
			if ans != expected[i] {
				t.Errorf("Received: %d; Expected %d", ans, expected[i])
			}
		})
	}
}

func TestPart2(t *testing.T) {
	var expected = []int{19, 23, 23, 29, 26}

	for i, line := range strings.Split(string(test), "\n") {
		testname := fmt.Sprintf("Line %d", i)
		t.Run(testname, func(t *testing.T) {
			ans := runPart2([]byte(line))
			if ans != expected[i] {
				t.Errorf("Received: %d; Expected %d", ans, expected[i])
			}
		})
	}
}

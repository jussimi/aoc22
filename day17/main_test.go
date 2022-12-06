package main

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed test.txt
var test []byte

func TestModule(t *testing.T) {
	const expected = 24000 // REPLACE ME
	fmt.Println(test)
	ans := run(test)
	if ans != expected {
		t.Errorf("Received: %d; Expected %d", ans, expected)
	}
}

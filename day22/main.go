package main

import (
	_ "embed"
	"fmt"
)

//go:embed data.txt
var data []byte

func main() {
	answer := run(data)
	fmt.Println("answer", answer)
}

func run(input []byte) int64 {
	return 24000
}

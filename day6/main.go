package main

import (
	_ "embed"
	"fmt"
)

//go:embed data.txt
var data []byte

func main() {
	fmt.Println("answer1", runPart1(data))
	fmt.Println("answer2", runPart2(data))
}

func areUnique(bytes []byte) bool {
	byteMap := make(map[byte]bool)
	for _, b := range bytes {
		if exists := byteMap[b]; exists {
			return false
		}
		byteMap[b] = true
	}
	return true
}

func findFirstUniqueBytes(input []byte, amount int) int {
	queue := []byte{}
	for i, b := range input {
		queue = append(queue, b)

		if len(queue) < amount+1 {
			continue
		}
		queue = queue[1:]

		if areUnique(queue) {
			return i + 1
		}
	}
	panic("not found")
}

func runPart1(input []byte) int {
	return findFirstUniqueBytes(input, 4)
}

func runPart2(input []byte) int {
	return findFirstUniqueBytes(input, 14)
}

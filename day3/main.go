package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed data.txt
var data []byte

func getPriorities() map[byte]int {
	alphabet := make([]byte, 0, 52)

	var ch byte
	for ch = 'a'; ch <= 'z'; ch++ {
		alphabet = append(alphabet, ch)
	}

	for ch = 'A'; ch <= 'Z'; ch++ {
		alphabet = append(alphabet, ch)
	}

	priorities := make(map[byte]int, 52)
	for i, ch := range alphabet {
		priorities[ch] = i + 1
	}

	return priorities
}

func main() {
	fmt.Println("answer1", runPart1(data))
	fmt.Println("answer2", runPart2(data))
}

func findDuplicateItem(rucksack string) byte {
	half := len(rucksack) / 2

	compartmentOne := make(map[rune]bool, half)

	for i, ch := range rucksack {
		if i < half {
			compartmentOne[ch] = true
		} else if compartmentOne[ch] {
			return byte(ch)
		}
	}
	panic("not found")
}

func runPart1(input []byte) int {
	priorities := getPriorities()

	sum := 0
	for _, rucksack := range strings.Split((string(input)), "\n") {
		duplicate := findDuplicateItem(rucksack)
		priority := priorities[duplicate]
		sum += priority
	}
	return sum
}

func runPart2(input []byte) int {
	priorities := getPriorities()

	badges := []rune{}

	charCounts := make(map[rune]int)
	for i, rucksack := range strings.Split((string(input)), "\n") {
		charsInRucksack := make(map[rune]bool)
		for _, ch := range rucksack {
			_, exists := charsInRucksack[ch]
			if !exists {
				charCounts[ch] += 1
				charsInRucksack[ch] = true
			}
		}
		if (i+1)%3 != 0 {
			continue
		}
		for k, n := range charCounts {
			if n == 3 {
				badges = append(badges, k)
				charCounts = make(map[rune]int)
				break
			}
		}
	}

	sum := 0
	for _, n := range badges {
		priority := priorities[byte(n)]
		sum += priority
	}
	return sum
}

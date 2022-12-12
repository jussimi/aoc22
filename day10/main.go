package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data.txt
var data []byte

func main() {
	fmt.Println("answer1", runPart1(data))
	fmt.Println("answer2")
	fmt.Println(runPart2(data))
}

func createHistory(input []byte) []int {
	x := 1

	history := []int{x}

	for _, instruction := range strings.Split(string(input), "\n") {
		split := strings.Fields(instruction)

		if split[0] == "addx" {
			amount, _ := strconv.Atoi(split[1])

			for _, val := range []int{0, amount} {
				x += val
				history = append(history, x)
			}
		} else {
			history = append(history, x)
		}
	}
	return history
}

func runPart1(input []byte) int {
	history := createHistory(input)
	indices := []int{20, 60, 100, 140, 180, 220}

	s := 0
	for _, i := range indices {
		s += history[i-1] * i
	}

	return s
}

func runPart2(input []byte) string {
	lines := []string{}

	history := createHistory(input)

	for i := 0; i < 6; i += 1 {
		line := make([]string, 40)
		for j := range line {
			line[j] = "."
		}

		for j := range line {
			x := history[0]
			history = history[1:]

			if j == x || j == x-1 || j == x+1 {
				line[j] = "#"
			}
		}
		lines = append(lines, strings.Join(line, ""))
	}

	return strings.Join(lines, "\n")
}

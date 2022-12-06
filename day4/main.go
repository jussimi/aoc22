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
	fmt.Println("answer2", runPart2(data))
}

type Interval struct {
	start int
	end   int
}

func (a Interval) isFullyOverlapping(b Interval) bool {
	return a.start <= b.start && a.end >= b.end
}

func (a Interval) isOverlapping(b Interval) bool {
	overlapStart := a.start <= b.start && a.end >= b.start
	overlapEnd := a.end >= b.end && b.end >= a.start
	return overlapStart || overlapEnd || areFullyOverlapping(a, b)
}

func areFullyOverlapping(a Interval, b Interval) bool {
	return a.isFullyOverlapping(b) || b.isFullyOverlapping(a)
}

func createInterval(str string) Interval {
	split := strings.Split(str, "-")
	start, _ := strconv.Atoi(split[0])
	end, _ := strconv.Atoi(split[1])
	return Interval{
		start,
		end,
	}
}

func getPairs(input []byte) [][2]Interval {
	pairs := [][2]Interval{}
	for _, line := range strings.Split((string(input)), "\n") {
		split := strings.Split(line, ",")
		pair := [2]Interval{
			createInterval(split[0]),
			createInterval(split[1]),
		}
		pairs = append(pairs, pair)
	}
	return pairs
}

func runPart1(input []byte) int {
	pairs := getPairs(input)

	sum := 0
	for _, pair := range pairs {
		if areFullyOverlapping(pair[0], pair[1]) {
			sum += 1
		}
	}
	return sum
}

func runPart2(input []byte) int {
	pairs := getPairs(input)

	sum := 0
	for _, pair := range pairs {
		if pair[0].isOverlapping(pair[1]) {
			sum += 1
		}
	}
	return sum
}

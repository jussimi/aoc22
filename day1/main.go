package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed data.txt
var data []byte

func main() {
	fmt.Println("answer1", runPart1(data))
	fmt.Println("answer2", runPart2(data))
}

func findElves(input []byte) [][]string {
	str := strings.Replace(string(input), "\r", "", -1)
	split := strings.Split(str, "\n\n")
	elves := [][]string{}
	for _, elfLines := range split {
		elves = append(elves, strings.Split(elfLines, "\n"))
	}
	return elves
}

func findElvesSumN(input []byte, n int) int {
	elves := findElves(input)

	sums := []int{}
	for _, elf := range elves {
		sum := 0
		for _, line := range elf {
			num, _ := strconv.Atoi(line)
			sum += num
		}
		sums = append(sums, sum)
	}

	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})

	sumOfMaxN := 0
	for _, s := range sums[0:n] {
		sumOfMaxN += s
	}
	return sumOfMaxN
}

func runPart1(input []byte) int {
	return findElvesSumN(input, 1)
}

func runPart2(input []byte) int {
	return findElvesSumN(input, 3)
}

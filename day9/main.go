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

var directions = map[string][2]int{
	"R": {0, 1},
	"L": {0, -1},
	"U": {1, 0},
	"D": {-1, 0},
}

func id(point [2]int) string {
	return fmt.Sprintf("%d-%d", point[0], point[1])
}

func sign(n int) int {
	if n < 0 {
		return -1
	}
	return 1
}

func abs(n int) int {
	return sign(n) * n
}

func moveRope(input []byte, knotCount int) int {

	knots := make([][2]int, knotCount)
	for i := range knots {
		knots[i] = [2]int{0, 0}
	}

	visited := make(map[string]bool)
	visited[id(knots[0])] = true

	for _, instruction := range strings.Split(string(input), "\n") {
		split := strings.Fields(instruction)
		direction := directions[split[0]]
		amount, _ := strconv.Atoi(split[1])

		for i := 0; i < amount; i += 1 {
			head := knots[0]
			knots[0] = [2]int{
				head[0] + direction[0],
				head[1] + direction[1],
			}

			for j := 1; j < knotCount; j += 1 {
				current := knots[j]
				previous := knots[j-1]
				dx := previous[0] - current[0]
				dy := previous[1] - current[1]

				if abs(dx) == 2 {
					current[0] = current[0] + sign(dx)
					if abs(dy) == 1 {
						current[1] = previous[1]
					}
				}

				if abs(dy) == 2 {
					current[1] = current[1] + sign(dy)
					if abs(dx) == 1 {
						current[0] = previous[0]
					}
				}
				knots[j] = current
			}

			tail := knots[len(knots)-1]
			visited[id(tail)] = true
		}
	}

	return len(visited)
}

func runPart1(input []byte) int {
	return moveRope(input, 2)
}

func runPart2(input []byte) int {
	return moveRope(input, 10)
}

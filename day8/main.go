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

type tree struct {
	size    int
	score   int
	visible bool
}

func createTreeGrid(input []byte) [][]*tree {
	grid := [][]*tree{}

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		row := []*tree{}
		for _, b := range line {
			n, _ := strconv.Atoi(string(b))
			row = append(row, &tree{size: n, visible: false, score: 1})
		}
		grid = append(grid, row)
	}
	return grid
}

func transposeGrid(grid [][]*tree) [][]*tree {
	n := len(grid[0])
	m := len(grid)

	transpose := make([][]*tree, n)

	for i := range transpose {
		transpose[i] = make([]*tree, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			transpose[i][j] = grid[j][i]
		}
	}
	return transpose
}

func reverseRow(row []*tree) []*tree {
	n := len(row)
	reversed := make([]*tree, n)

	for i := 0; i < n; i += 1 {
		reversed[n-i-1] = row[i]
	}
	return reversed
}

func findDistance(row []*tree, idx int) int {
	if idx == 0 {
		return 0
	}
	distance := 1
	for distance < idx {
		if row[idx].size <= row[idx-distance].size {
			break
		}
		distance += 1
	}
	return distance
}

func findVisibles(row []*tree) {
	for j, tree := range row {
		distance := findDistance(row, j)
		tree.score *= distance
		if j == 0 || (distance == j && tree.size > row[0].size) {
			tree.visible = true
		}
	}
}

func loopRows(grid [][]*tree) {
	for _, row := range grid {
		findVisibles(row)
		findVisibles(reverseRow(row))
	}
}

func runPart1(input []byte) int {

	grid := createTreeGrid(input)
	transpose := transposeGrid(grid)

	loopRows(grid)
	loopRows(transpose)

	count := 0
	for _, row := range grid {
		for _, tree := range row {
			if tree.visible {
				count += 1
			}
		}
	}

	return count
}

func runPart2(input []byte) int {
	grid := createTreeGrid(input)
	transpose := transposeGrid(grid)

	loopRows(grid)
	loopRows(transpose)

	largest := 0
	for _, row := range grid {
		for _, tree := range row {
			if tree.score > largest {
				largest = tree.score
			}
		}
	}

	return largest
}

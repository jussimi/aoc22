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

type Instruction struct {
	from   int
	to     int
	amount int
}

func parsePiles(pileStr string) [][]string {
	lines := strings.Split(pileStr, "\n")

	col_labels := strings.Fields(lines[len(lines)-1])
	col_amount := len(col_labels)

	cols := make([][]string, col_amount)

	for i := range cols {
		cols[i] = []string{}
	}

	for i := len(lines) - 2; i >= 0; i -= 1 {
		for j := 0; j < col_amount; j += 1 {
			val := string(lines[i][4*j+1])
			if val != " " {
				cols[j] = append(cols[j], val)
			}
		}
	}

	return cols
}

func parseInstructions(instructionsStr string) []Instruction {
	replaced := strings.Replace(instructionsStr, "move ", "", -1)
	replaced = strings.Replace(replaced, " from ", ",", -1)
	replaced = strings.Replace(replaced, " to ", ",", -1)

	instructions := []Instruction{}
	for _, inst := range strings.Split(replaced, "\n") {
		split := strings.Split(inst, ",")
		amount, _ := strconv.Atoi(split[0])
		from, _ := strconv.Atoi(split[1])
		to, _ := strconv.Atoi(split[2])
		instructions = append(instructions, Instruction{
			from:   from - 1,
			to:     to - 1,
			amount: amount,
		})
	}

	return instructions
}

func runPart1(input []byte) string {
	split := strings.Split(string(input), "\n\n")

	piles := parsePiles(split[0])
	instructions := parseInstructions(split[1])

	for _, instruction := range instructions {
		for i := 0; i < instruction.amount; i += 1 {
			from := piles[instruction.from]
			to := piles[instruction.to]
			toMove := from[len(from)-1]
			piles[instruction.to] = append(to, toMove)
			piles[instruction.from] = from[:len(from)-1]
		}
	}

	str := ""
	for _, pile := range piles {
		str += pile[len(pile)-1]
	}

	return str
}

func runPart2(input []byte) string {
	split := strings.Split(string(input), "\n\n")
	piles := parsePiles(split[0])
	instructions := parseInstructions(split[1])

	for _, instruction := range instructions {
		from := piles[instruction.from]
		to := piles[instruction.to]
		toMove := from[len(from)-instruction.amount:]
		piles[instruction.to] = append(to, toMove...)
		piles[instruction.from] = from[:len(from)-instruction.amount]
	}

	str := ""
	for _, pile := range piles {
		str += pile[len(pile)-1]
	}

	return str
}

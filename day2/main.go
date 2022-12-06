package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed data.txt
var data []byte

func main() {
	fmt.Println("answer1", runPart1(data))
	fmt.Println("answer2", runPart2(data))
}

type RPS = int32

const (
	Rock RPS = iota + 1 // Also represents score.
	Paper
	Scissors
)

const Win = 6
const Loss = 0
const Draw = 3

var opponentMap = map[string]RPS{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}

var playerMap = map[string]RPS{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var resultMap = map[string]int32{
	"X": Loss,
	"Y": Draw,
	"Z": Win,
}

var winMap = map[RPS]RPS{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

var lossMap = map[RPS]RPS{
	Rock:     Paper,
	Paper:    Scissors,
	Scissors: Rock,
}

func getScore(player RPS, opponent RPS) int32 {
	if player == opponent {
		return player + Draw
	}
	if winMap[player] == opponent {
		return player + Win
	}
	return player
}

func runPart1(input []byte) int32 {
	var score int32 = 0

	for _, line := range strings.Split((string(input)), "\n") {
		moves := strings.Fields(line)
		opponentMove := opponentMap[moves[0]]
		playerMove := playerMap[moves[1]]
		score += getScore(playerMove, opponentMove)
	}

	return score
}

func runPart2(input []byte) int32 {
	var score int32 = 0

	for _, line := range strings.Split((string(input)), "\n") {
		moves := strings.Fields(line)
		opponentMove := opponentMap[moves[0]]

		result := resultMap[moves[1]]

		var playerMove RPS
		switch result {
		case Draw:
			playerMove = opponentMove
		case Win:
			playerMove = lossMap[opponentMove]
		case Loss:
			playerMove = winMap[opponentMove]
		}

		score += getScore(playerMove, opponentMove)
	}

	return score
}

package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	OPPONENT_ROCK     string = "A"
	OPPONENT_PAPER           = "B"
	OPPONENT_SCISSORS        = "C"
)

const (
	ELF_ROCK     string = "X"
	ELF_PAPER           = "Y"
	ELF_SCISSORS        = "Z"
)

const (
	SHOULD_LOSE string = "X"
	SHOULD_DRAW        = "Y"
	SHOULD_WIN         = "Z"
)

func main() {
	bytes, _ := os.ReadFile("2022/day2.txt")
	puzzleInput := string(bytes)
	lines := strings.Split(puzzleInput, "\n")

	score := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		opponentsHand := parts[0]
		secondSymbol := parts[1]

		elfsHand := secondSymbol

		// comment out for part two
		elfsHand = getElfsHand(opponentsHand, secondSymbol)

		switch elfsHand {
		case ELF_ROCK:
			score += 1
		case ELF_PAPER:
			score += 2
		case ELF_SCISSORS:
			score += 3
		}

		score += getRoundScore(opponentsHand, elfsHand)
	}

	fmt.Println("Answer ")
	fmt.Println(score)
}

func getRoundScore(opponentHand string, elfHand string) int {
	if opponentHand == "A" { //rock
		switch elfHand {
		case "X": // rock
			return 3
		case "Y": // paper
			return 6
		case "Z": // scissors
			return 0
		}
	}

	if opponentHand == "B" { //paper
		switch elfHand {
		case "X": // rock
			return 0
		case "Y": // paper
			return 3
		case "Z": // scissors
			return 6
		}
	}

	if opponentHand == "C" { //scissors
		switch elfHand {
		case "X": // rock
			return 6
		case "Y": // paper
			return 0
		case "Z": // scissors
			return 3
		}
	}

	return 0
}

func getElfsHand(opponentHand string, winLoseDraw string) string {
	if opponentHand == OPPONENT_ROCK {
		switch winLoseDraw {
		case SHOULD_LOSE:
			return ELF_SCISSORS
		case SHOULD_DRAW:
			return ELF_ROCK
		case SHOULD_WIN:
			return ELF_PAPER
		}
	}

	if opponentHand == OPPONENT_PAPER {
		switch winLoseDraw {
		case SHOULD_LOSE:
			return ELF_ROCK
		case SHOULD_DRAW:
			return ELF_PAPER
		case SHOULD_WIN:
			return ELF_SCISSORS
		}
	}

	if opponentHand == OPPONENT_SCISSORS {
		switch winLoseDraw {
		case SHOULD_LOSE:
			return ELF_PAPER
		case SHOULD_DRAW:
			return ELF_SCISSORS
		case SHOULD_WIN:
			return ELF_ROCK
		}
	}

	return ""
}

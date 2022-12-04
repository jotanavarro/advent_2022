package main

import (
	"advent_2022/internal/pkg/filemodel"
	"advent_2022/internal/pkg/helper"
	"fmt"
	"log"
	"strings"
)

// https://adventofcode.com/2022/day/1

func main() {
	scanner := filemodel.CreateScanner("02", "01")
	defer func(scanner filemodel.FileScanner) {
		err := scanner.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(scanner)

	firstPoints := 0
	secondPoints := 0

	for scanner.Scan() {
		play := strings.Fields(scanner.Text())

		firstPoints += evaluate(translate(play[0]), translate(play[1]))
		secondPoints += calculateYourPlay(translate(play[0]), translate(play[1]))
	}

	fmt.Printf("1) The total points are: %d\n", firstPoints)
	fmt.Printf("2) The total points are: %d\n", secondPoints)
}

/*
Translates a play into integers.
*/
func translate(choice string) (result int) {
	if choice == "A" || choice == "X" {
		result = 0
	} else if choice == "B" || choice == "Y" {
		result = 1
	} else {
		result = 2
	}

	return result
}

/*
How to evaluate a play for the first exercise.
*/
func evaluate(opponent, you int) (result int) {
	result += you
	if opponent == you {
		result += 3
	} else if (opponent == 0 && you == 1) || (opponent == 1 && you == 2) || (opponent == 2 && you == 0) {
		result += 6
	}

	return result + 1
}

/*
How to evaluate a play for the second exercise.
*/
func calculateYourPlay(opponent, you int) (result int) {
	var play int
	if you == 0 {
		// You have to lose
		play = helper.MyMod(opponent-1, 3)
	} else if you == 1 {
		// You have to reach a draw
		play = opponent
	} else {
		// You have to win
		play = helper.MyMod(opponent+1, 3)
	}

	return evaluate(opponent, play)
}

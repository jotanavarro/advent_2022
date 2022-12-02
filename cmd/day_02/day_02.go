package main

import (
	"advent_2022/internal/pkg/input_handler"
	"fmt"
	"log"
	"strings"
)

// https://adventofcode.com/2022/day/1

func main() {
	scanner := input_handler.CreateScanner("02", "01")
	defer func(scanner input_handler.FileScanner) {
		err := scanner.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(scanner)

	points := 0

	for scanner.Scan() {
		play := strings.Fields(scanner.Text())

		points += evaluate(translate(play[0]), translate(play[1]))
	}

	fmt.Printf("The total points are: %d\n", points)
}

func translate(choice string) (result int) {
	if choice == "A" || choice == "X" {
		result = 1
	} else if choice == "B" || choice == "Y" {
		result = 2
	} else {
		result = 3
	}

	return result
}

func evaluate(opponent, you int) (result int) {
	result += you
	if opponent == you {
		result += 3
	} else if (opponent == 1 && you == 2) || (opponent == 2 && you == 3) || (opponent == 3 && you == 1) {
		result += 6
	}

	return result
}

package main

import (
	"advent_2022/internal/pkg/filemodel"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
)

// https://adventofcode.com/2022/day/3

func main() {
	scanner := filemodel.CreateScanner("03", "01")
	defer func(scanner filemodel.FileScanner) {
		err := scanner.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(scanner)

	calculatePriority(scanner)
}

func calculatePriority(scanner filemodel.FileScanner) {
	totalPriority := 0

	for scanner.Scan() {
		duplicate, err := findDuplicate(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		priority, _ := calculateScore(duplicate)

		totalPriority += priority
	}

	fmt.Printf("Total priority for the sum of all backpacks is: %d.\n", totalPriority)
}

func findDuplicate(content string) (duplicate string, err error) {
	midpoint := len(content) / 2

	// We will be naive and ignore the potential duplicates which are not present in the second half.
	for pos, char := range content {
		if pos > midpoint {
			break
		}

		if strings.ContainsRune(content[midpoint:], char) {
			return string(char), nil
		}
	}

	return "", errors.New(fmt.Sprintf("No duplicate found in %s!", content))
}

/*
calculateScore will use the position of a letter in the alphabet to calculate the score of a letter.  In case we pass
more than a single letter an error will be returned.
*/
func calculateScore(letter string) (score int, err error) {
	numLetters := len(letter)
	if numLetters != 1 {
		return -1, errors.New(fmt.Sprintf("Passed a string of length %d instead of 1!", numLetters))
	}

	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	r, _ := regexp.Compile(letter)

	return r.FindStringIndex(alphabet)[0] + 1, nil
}

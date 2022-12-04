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

	scanner = filemodel.CreateScanner("03", "01")
	defer func(scanner filemodel.FileScanner) {
		err := scanner.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(scanner)

	calculateBadgePriorities(scanner)
}

/*
calculateBadgePriorities will solve the second part of the puzzle.
*/
func calculateBadgePriorities(scanner filemodel.FileScanner) {
	badgePriority := 0

	for scanner.Scan() {
		var elfGroup []string
		elfGroup = append(elfGroup, scanner.Text())

		for i := 0; i < 2; i++ {
			scanner.Scan()
			elfGroup = append(elfGroup, scanner.Text())
		}

		badge, err := findBadge(elfGroup)
		if err != nil {
			log.Fatal(err)
		}

		badgeScore, err := calculateScore(badge)
		if err != nil {
			log.Fatal(err)
		}

		badgePriority += badgeScore
	}

	fmt.Printf("The total priority of the badges is: %d\n", badgePriority)
}

/*
findBadge will find in a triad of strings the first letter present in all of them, then return it.
*/
func findBadge(elfGroup []string) (badge string, err error) {
	for _, char := range elfGroup[0] {
		if strings.ContainsRune(elfGroup[1], char) && strings.ContainsRune(elfGroup[2], char) {
			return string(char), nil
		}
	}
	return "", errors.New("No duplicate found!")
}

/*
calculatePriority will solve the first part of the puzzle.
*/
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

/*
findDuplicate will find any character in the first half of a string which is present in the second half.  This function
is very naive assuming all the strings are of an even length and that we will only return the first duplicate, ignoring
any other potential duplicate.
*/
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

package main

import (
	"advent_2022/internal/pkg/filemodel"
	"errors"
	"fmt"
	"log"
)

// https://adventofcode.com/2022/day/6

func main() {
	scanner := filemodel.CreateScanner("06", "01")
	defer func(scanner filemodel.FileScanner) {
		err := scanner.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(scanner)

	for scanner.Scan() {
		signal := scanner.Text()
		marker, err := processSignal(signal, 4)
		if err != nil {
			fmt.Printf("No marker found in '%s'.\n", signal)
		}
		message, err := processSignal(signal, 14)
		if err != nil {
			fmt.Printf("No message found in '%s'.\n", signal)
		}

		fmt.Printf("Marker found at position: '%d'.\n", marker)
		fmt.Printf("Message found at position: '%d'.\n", message)

	}
}

func processSignal(signal string, length int) (position int, err error) {
	buffer := ""

	for k, _ := range signal {
		if k > length-1 {
			buffer = signal[k-length : k]
			if !hasDuplicates(buffer) {
				return k, nil
			}
		}
	}

	return -1, errors.New("no marker found")
}

func hasDuplicates(signal string) bool {
	index := make(map[rune]bool)

	for _, letter := range signal {
		if index[letter] {
			return true
		} else {
			index[letter] = true
		}
	}

	return false
}

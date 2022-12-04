package main

import (
	"advent_2022/internal/pkg/filemodel"
	"fmt"
	"log"
	"strconv"
)

// https://adventofcode.com/2022/day/1

func main() {
	scanner := filemodel.CreateScanner("01", "01")
	defer func(scanner filemodel.FileScanner) {
		err := scanner.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(scanner)

	first, second, third := 0, 0, 0
	calories := 0

	for scanner.Scan() {
		input := scanner.Text()

		if input == "" {
			if calories > third {
				if calories > second {
					third = second
					if calories > first {
						second = first
						first = calories
					} else {
						second = calories
					}
				} else {
					third = calories
				}
			}
			calories = 0
		} else {
			number, err := strconv.Atoi(input)
			if err != nil {
				log.Fatal(err)
			}

			calories += number
		}
	}

	fmt.Printf("The highest amount of calories is:\n%dKcal\n", first)
	fmt.Printf("The total calores carried by the top three is:\n%dKcal\n", first+second+third)
}

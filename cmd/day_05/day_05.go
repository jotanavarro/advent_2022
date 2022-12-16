package main

import (
	"advent_2022/internal/pkg/filemodel"
	"advent_2022/internal/pkg/helper"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/5

func main() {
	scanner := filemodel.CreateScanner("05", "01")
	defer func(scanner filemodel.FileScanner) {
		err := scanner.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(scanner)

	crateMap := loadCrateMap(scanner)

	for scanner.Scan() {
		request := strings.Split(scanner.Text(), " ")
		crateNumber, _ := strconv.Atoi(request[1])
		from, _ := strconv.Atoi(request[3])
		to, _ := strconv.Atoi(request[5])
		//moveCrates(crateMap, crateNumber, from, to)
		moveCratesTogether(crateMap, crateNumber, from, to)
	}

	println(findTopCrates(crateMap))
}

func loadCrateMap(scanner filemodel.FileScanner) (crates map[int]*helper.List) {
	crates = make(map[int]*helper.List)

	for scanner.Scan() {
		input := scanner.Text()

		if input == "" {
			break
		} else if input[1] != '1' {
			// Process the crate position
			column := 1
			position := 1
			for position < len(input) {
				value := string(input[position])

				if value != " " {
					if crates[column] == nil {
						tmpList := helper.List{}
						tmpList.InsertReversed(value)

						crates[column] = &tmpList
					} else {
						crates[column].InsertReversed(value)
					}
				}

				position += 4
				column++
			}

		}
	}

	return
}

func findTopCrates(crateMap map[int]*helper.List) (topCrates string) {
	columns := len(crateMap)

	for i := 1; i <= columns; i++ {
		topCrates += fmt.Sprintf("%s", (crateMap[i]).Pop())
	}

	return
}

func moveCrates(crateMap map[int]*helper.List, crateNumber, origin, destination int) {
	for i := 0; i < crateNumber; i++ {
		crate := (crateMap[origin]).Pop()
		(crateMap[destination]).Insert(crate)
	}
}

func moveCratesTogether(crateMap map[int]*helper.List, crateNumber, origin, destination int) {
	var crates []string

	for i := 0; i < crateNumber; i++ {
		crate := (crateMap[origin]).Pop()
		crates = append(crates, fmt.Sprintf("%s", crate))
	}

	for i := len(crates) - 1; i >= 0; i-- {
		(crateMap[destination]).Insert(crates[i])
	}
}

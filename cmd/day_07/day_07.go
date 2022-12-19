package main

import (
	"advent_2022/internal/pkg/filemodel"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/7

func main() {
	scanner := filemodel.CreateScanner("07", "01")
	defer func(scanner filemodel.FileScanner) {
		err := scanner.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(scanner)

	fileSystem := loadFileSystem(scanner)

	// Exercise 1
	var smallerDirectories []item
	var totalSize int
	fileSystem.findDirectoriesSmallerThan(100000, &smallerDirectories)
	for _, v := range smallerDirectories {
		totalSize += v.size
	}
	fmt.Printf("Total size: %d\n", totalSize)

	// Exercise 2
	availableSpace := 70000000 - fileSystem.size
	requiredSpace := 30000000
	toDelete := requiredSpace - availableSpace

	var goodForDeletion []item
	smallest := item{size: math.MaxInt}

	fileSystem.findDirectoriesGreaterThan(toDelete, &goodForDeletion)

	for _, v := range goodForDeletion {
		if v.size < smallest.size {
			smallest = v
		}
	}

	fmt.Printf("\n%d", toDelete-smallest.size)
}

func loadFileSystem(scanner filemodel.FileScanner) *directory {
	fileSystem := directory{parent: nil, name: "/", size: 0, directories: map[string]*directory{}}

	currentDirectory := &fileSystem
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()

		parameters := strings.Split(line, " ")

		if parameters[0] == "$" {
			if parameters[1] == "cd" {
				targetDirectory := parameters[2]
				tmpDir, err := currentDirectory.crawlDirectory(targetDirectory)
				if err != nil {
					log.Fatal(fmt.Sprintf("error: %s directory does not exist.", targetDirectory))
				}
				currentDirectory = tmpDir
			}
		} else if parameters[0] == "dir" {
			currentDirectory.createDirectory(parameters[1])
		} else {
			size, _ := strconv.Atoi(parameters[0])
			currentDirectory.createFile(parameters[1], size)
		}
	}

	return &fileSystem
}

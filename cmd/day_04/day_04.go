package main

import (
	"advent_2022/internal/pkg/filemodel"
	"advent_2022/internal/pkg/helper"
	"fmt"
	"log"
	"strings"
)

// https://adventofcode.com/2022/day/4

func main() {
	scanner := filemodel.CreateScanner("04", "01")
	defer func(scanner filemodel.FileScanner) {
		err := scanner.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(scanner)

	completeOverlap := 0
	partialOverlap := 0

	for scanner.Scan() {
		elves := strings.Split(scanner.Text(), ",")

		left, err := helper.StringSliceToInt(strings.Split(elves[0], "-"))
		if err != nil {
			log.Fatal(err)
		}
		right, err := helper.StringSliceToInt(strings.Split(elves[1], "-"))
		if err != nil {
			log.Fatal(err)
		}

		leftLength, rightLength := calculateLength(left), calculateLength(right)

		if leftLength >= rightLength {
			if findCompleteOverlap(right, left) {
				completeOverlap++
			}
		} else {
			if findCompleteOverlap(left, right) {
				completeOverlap++
			}
		}

		if findPartialOverlap(left, right) {
			partialOverlap++
		}
	}

	fmt.Printf("The number of completely overlapped sections is: %d.\n", completeOverlap)
	fmt.Printf("The number of partially overlapped sections is: %d.\n", partialOverlap)
}

/*
calculateLength will assume that a section consisting of two integers as strings, where the first one is smaller than
the second one if passed.
*/
func calculateLength(sections []int) int {
	return sections[1] - sections[0]
}

/*
findCompleteOverlap assumes that there is an overlap between sections if the beginning of the smaller section is equal or
greater than the one for the larger section and also if the end of the smaller section is equal or smaller than the one
of the larger section.
*/
func findCompleteOverlap(smaller, larger []int) (overlap bool) {
	if smaller[0] >= larger[0] && smaller[1] <= larger[1] {
		return true
	} else {
		return false
	}
}

/*
findPartialOverlap will tell if there is any point of any section within the other one, that necessarily implies there
is at least one overlap.
*/
func findPartialOverlap(left, right []int) (overlap bool) {
	a, b := left[0], left[1]
	c, d := right[0], right[1]

	if ((c <= a) && (a <= d)) || ((c <= b) && (b <= d)) {
		overlap = true
	} else if ((a <= c) && (c <= b)) || ((a <= d) && (d <= b)) {
		overlap = true
	}

	return
}

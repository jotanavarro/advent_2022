package helper

import (
	"log"
	"strconv"
)

/*
StringSliceToInt will take a slice of strings and convert each element to an integer.
*/
func StringSliceToInt(input []string) (output []int, err error) {
	for _, char := range input {
		number, err := strconv.Atoi(char)
		if err != nil {
			log.Fatal(err)
		}

		output = append(output, number)
	}

	return
}

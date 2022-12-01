package input_handler

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func CreateScanner(day string, exercise string) FileScanner {
	filename := fmt.Sprintf("assets/input/day_%s_exercise_%s.txt", day, exercise)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Unable to find/open '%s'!", filename)
	}

	return FileScanner{file, bufio.NewScanner(file)}
}

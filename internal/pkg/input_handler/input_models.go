package input_handler

import (
	"bufio"
	"io"
)

type FileScanner struct {
	io.Closer
	*bufio.Scanner
}

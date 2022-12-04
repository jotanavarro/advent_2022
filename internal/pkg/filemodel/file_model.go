package filemodel

import (
	"bufio"
	"io"
)

type FileScanner struct {
	io.Closer
	*bufio.Scanner
}

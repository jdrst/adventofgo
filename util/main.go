package util

import (
	"log"
	"os"
	"runtime"
	"strings"
)

func newLine() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}

//File is a file as bytearray
type File []byte

//ReadFile reads the file into
func ReadFile(path string) File {
	input, err := os.ReadFile(path)
	Handle(err)
	return input
}

//Handle calls log.Fatal on an Error
func Handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Returns the lines of a file as string array
func (f File) AsLines() []string {
	return strings.Split(strings.TrimSpace(string(f)), newLine())
}

//Returns the File with CRLF linebreaks instead of LF (for testing purposes)
func (f File) WithCRLF() File {
	return []byte(strings.ReplaceAll(string(f), "\n", newLine()))
}

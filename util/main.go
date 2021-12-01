package util

import (
	"log"
	"os"
	"runtime"
	"strconv"
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

//Lines is an array of Line
type Lines []Line

//Line is a string
type Line string

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

//Returns the lines of a file as Lines type (string array)
func (f File) AsLines() Lines {
	strings := strings.Split(strings.TrimSpace(string(f)), newLine())
	lines := make([]Line, len(strings))
	for i, s := range strings {
		lines[i] = Line(s)
	}
	return lines
}

//Returns the File with CRLF linebreaks instead of LF (for testing purposes)
func (f File) WithCRLF() File {
	return []byte(strings.ReplaceAll(string(f), "\n", newLine()))
}

//Returns lines as ints
func (lines Lines) AsInts() []int {
	res := make([]int, len(lines))
	for i, l := range lines {
		res[i] = l.AsInt()
	}
	return res
}

//Returns the line converted to int
func (l Line) AsInt() int {
	i, err := strconv.Atoi(string(l))
	Handle(err)
	return i
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	Handle(err)
	return i
}

package util

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

//NewLine returns "\n" for non-windows and "\r\n" for windows runtime
func NewLine() string {
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

type Point struct {
	X, Y int
}

//ReadFile reads a file into the File struct ([]byte) and calls log.Fatal on an error
func ReadFile(path string) File {
	input, err := os.ReadFile(path)
	Handle(err)
	return input
}

//Handle calls log.Fatal on an error
func Handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//AsLines returns the lines of a file as Lines type ([]string)
func (f File) AsLines() Lines {
	return split(strings.TrimSpace(string(f)), NewLine())
}

//SubSplitWith splits a Line on into Lines using the given separator
func (l Line) SubSplitWith(separator string) Lines {
	return split(string(l), separator)
}

func split(s, sep string) Lines {
	strings := strings.Split(s, sep)
	lines := make([]Line, len(strings))
	for i, s := range strings {
		lines[i] = Line(s)
	}
	return lines
}

//WithOSLinebreaks returns the File with OS-specific linebreaks instead of LF (for testing purposes)
func (f File) WithOSLinebreaks() File {
	return []byte(strings.ReplaceAll(string(f), "\n", NewLine()))
}

//AsInts converts Lines ([]string) to []int
func (lines Lines) AsInts() []int {
	res := make([]int, len(lines))
	for i, l := range lines {
		res[i] = l.AsInt()
	}
	return res
}

//As2DInts converts Lines ([]string) to [][]int, splitting the line on the given separator
func (lines Lines) As2DInts(separator string) [][]int {
	res := make([][]int, len(lines))

	for i, l := range lines {
		res[i] = l.SubSplitWith(separator).AsInts()
	}
	return res
}

//AsInt converts a line to int and calls log.Fatal on an error
func (l Line) AsInt() int {
	return ToInt(string(l))
}

//ToInt converts a string to an int and calls log.Fatal on an error
func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	Handle(err)
	return i
}

//Neighbours returns direct neighbours on x and y axis (no diagonal neighbours)
func (n *Point) Neighbours(maxX, maxY int) []Point {
	res := make([]Point, 0)
	if n.X > 0 {
		res = append(res, Point{n.X - 1, n.Y})
	}
	if n.X < maxX {
		res = append(res, Point{n.X + 1, n.Y})
	}
	if n.Y > 0 {
		res = append(res, Point{n.X, n.Y - 1})
	}
	if n.Y < maxY {
		res = append(res, Point{n.X, n.Y + 1})
	}
	return res
}

//Neighbours returns direct neighbours on x and y axis as well as diagonal neighbours
func (n *Point) NeighboursWithDiagonal(maxX, maxY int) []Point {
	res := make([]Point, 0)
	if n.X > 0 {
		res = append(res, Point{n.X - 1, n.Y})
		if n.Y < maxY {
			res = append(res, Point{n.X - 1, n.Y + 1})
		}
	}
	if n.X < maxX {
		res = append(res, Point{n.X + 1, n.Y})
		if n.Y > 0 {
			res = append(res, Point{n.X + 1, n.Y - 1})
		}
	}
	if n.Y > 0 {
		res = append(res, Point{n.X, n.Y - 1})
		if n.X > 0 {
			res = append(res, Point{n.X - 1, n.Y - 1})
		}
	}
	if n.Y < maxY {
		res = append(res, Point{n.X, n.Y + 1})
		if n.X < maxX {
			res = append(res, Point{n.X + 1, n.Y + 1})
		}
	}
	return res
}

//ManhattanDistance returns the manhattan distance between two points (delta between x's plus delta between y's)
func ManhattanDistance(a, b Point) int {
	return delta(a.X, b.X) + delta(a.Y, b.Y)
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
func delta(x, y int) int {
	return abs(x - y)
}

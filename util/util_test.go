package util

import (
	"io/fs"
	"os"
	"reflect"
	"testing"
)

var testInput = `test
test2`

func TestReadFile(t *testing.T) {
	fileName := "test.txt"
	err := os.WriteFile(fileName, []byte(testInput), fs.ModeTemporary)
	Handle(err)
	defer os.Remove(fileName)
	expected := testInput

	actual := string(ReadFile(fileName))

	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

func TestAsLines(t *testing.T) {
	expected := Lines{"test", "test2"}

	actual := File(testInput).WithOSLinebreaks().AsLines()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

var subSplitTestCases = []struct {
	line      Line
	expected  Lines
	separator string
}{
	{Line("test,test2"), Lines{"test", "test2"}, ","},
	{Line("abc;;def"), Lines{"abc", "def"}, ";;"},
	{Line("test\ntest2"), Lines{"test", "test2"}, "\n"},
	{Line("test\x33Xestest2"), Lines{"test", "test2"}, "\x33Xes"},
}

func TestSubSplitWith(t *testing.T) {
	for _, test := range subSplitTestCases {
		actual := test.line.SubSplitWith(test.separator)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("split %v with %v failed. expected was: %v \n actual is: %v", test.line, test.separator, test.expected, actual)
		}
	}
}

func TestAsInts(t *testing.T) {
	lines := Lines{"1", "2", "3", "4"}
	expected := []int{1, 2, 3, 4}

	actual := lines.AsInts()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

func TestAsInt(t *testing.T) {
	lines := Line("42")
	expected := 42

	actual := lines.AsInt()

	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

func TestToInt(t *testing.T) {
	expected := 42
	actual := ToInt("42")

	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

func TestAs2DInts(t *testing.T) {
	lines := Lines{"1,2", "2,3", "3,4,5", "4,5"}
	expected := [][]int{{1, 2}, {2, 3}, {3, 4, 5}, {4, 5}}

	actual := lines.As2DInts(",")

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

var NeighbourTestCases = []struct {
	point      Point
	maxX, maxY int
	expected   []Point
}{
	{Point{1, 1}, 2, 2, []Point{{0, 1}, {2, 1}, {1, 0}, {1, 2}}},
	{Point{1, 1}, 1, 1, []Point{{0, 1}, {1, 0}}},
	{Point{0, 0}, 2, 2, []Point{{1, 0}, {0, 1}}},
	{Point{0, 4}, 2, 5, []Point{{1, 4}, {0, 3}, {0, 5}}},
}

func TestNeighbours(t *testing.T) {
	for _, test := range NeighbourTestCases {
		actual := test.point.Neighbours(test.maxX, test.maxY)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("expected was: %v \n actual is: %v", test.expected, actual)
		}
	}
}

var NeighboursWithDiagonalTestCases = []struct {
	point      Point
	maxX, maxY int
	expected   []Point
}{
	{Point{1, 1}, 2, 2, []Point{{0, 1}, {0, 2}, {2, 1}, {2, 0}, {1, 0}, {0, 0}, {1, 2}, {2, 2}}},
	{Point{1, 1}, 1, 1, []Point{{0, 1}, {1, 0}, {0, 0}}},
	{Point{0, 0}, 2, 2, []Point{{1, 0}, {0, 1}, {1, 1}}},
	{Point{0, 4}, 2, 5, []Point{{1, 4}, {1, 3}, {0, 3}, {0, 5}, {1, 5}}},
}

func TestNeighboursWithDiagonal(t *testing.T) {
	for _, test := range NeighboursWithDiagonalTestCases {
		actual := test.point.NeighboursWithDiagonal(test.maxX, test.maxY)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("expected was: %v \n actual is: %v", test.expected, actual)
		}
	}
}

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
	os.WriteFile(fileName, []byte(testInput), fs.ModeTemporary)
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

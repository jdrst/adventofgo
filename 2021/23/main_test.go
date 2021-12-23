package main

import (
	"testing"

	"github.com/jdrst/adventofgo/util"
)

var testInput = `#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`

var testInput2 = `#############
#...........#
###A#B#C#D###
  #A#B#C#D#
  #########`

var testInput3 = `#############
#...........#
###B#C#B#D###
  #D#C#B#A#
  #D#B#A#C#
  #A#D#C#A#
  #########`

var testInput4 = `#############
#...........#
###B#A#C#D###
  #A#B#C#D#
  #A#B#C#D#
  #A#B#C#D#
  #########`

func TestPartOne(t *testing.T) {
	expected := 12521
	actual := partOne(util.File(testInput).WithOSLinebreaks())
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}

	expected = 0
	actual = partOne(util.File(testInput2).WithOSLinebreaks())
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}

	expected = 14350
	actual = partOne(util.ReadFile("input.txt"))
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

func TestString(t *testing.T) {
	expected := "used energy: 0\n" + testInput
	// expected = testInput
	actual := makeBurrow(util.File(testInput).WithOSLinebreaks().AsLines()).String()
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}

	expected = "used energy: 0\n" + testInput3
	// expected = testInput3
	actual = makeBurrow(util.File(testInput3).WithOSLinebreaks().AsLines()).String()
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}

	expected = "used energy: 0\n" + testInput4
	// expected = testInput4
	actual = makeBurrow(util.File(testInput4).WithOSLinebreaks().AsLines()).String()
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 44169
	actual := partTwo(util.File(testInput).WithOSLinebreaks())
	if actual != expected {
		t.Errorf("\nexpected was: %v\nactual is: %v", expected, actual)
	}

	expected = 49742
	actual = partTwo(util.ReadFile("input.txt"))
	if actual != expected {
		t.Errorf("\nexpected was: %v\nactual is: %v", expected, actual)
	}
}

func BenchmarkPartOne(b *testing.B) {
	input := util.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		partOne(input)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	input := util.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		partTwo(input)
	}
}

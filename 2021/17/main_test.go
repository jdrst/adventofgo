package main

import (
	"testing"

	"github.com/jdrst/adventofgo/util"
)

var testInput = `target area: x=20..30, y=-10..-5`

func TestPartOne(t *testing.T) {
	expected := 45
	actual := partOne(util.File(testInput).WithOSLinebreaks())
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}

	expected = 2850
	actual = partOne(util.ReadFile("input.txt"))
	if actual != expected {
		t.Errorf("\nexpected was: %v\nactual is: %v", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 112
	actual := partTwo(util.File(testInput).WithOSLinebreaks())
	if actual != expected {
		t.Errorf("\nexpected was: %v\nactual is: %v", expected, actual)
	}

	expected = 1117
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

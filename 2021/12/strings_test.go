package main

import (
	"testing"

	"github.com/jdrst/adventofgo/util"
)

func TestPartOneWithUniquePathArray(t *testing.T) {
	expected := 10
	actual := partOneWithUniquePathArray(util.File(testInput).WithOSLinebreaks())
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}

	expected = 19
	actual = partOneWithUniquePathArray(util.File(testInput2).WithOSLinebreaks())
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}

	expected = 226
	actual = partOneWithUniquePathArray(util.File(testInput3).WithOSLinebreaks())
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

func TestPartTwoWithUniquePathArray(t *testing.T) {
	expected := 36
	actual := partTwoWithUniquePathArray(util.File(testInput).WithOSLinebreaks())
	if actual != expected {
		t.Errorf("\nexpected was: %v\nactual is: %v", expected, actual)
	}

	expected = 103
	actual = partTwoWithUniquePathArray(util.File(testInput2).WithOSLinebreaks())
	if actual != expected {
		t.Errorf("\nexpected was: %v\nactual is: %v", expected, actual)
	}

	expected = 3509
	actual = partTwoWithUniquePathArray(util.File(testInput3).WithOSLinebreaks())
	if actual != expected {
		t.Errorf("\nexpected was: %v\nactual is: %v", expected, actual)
	}
}

func BenchmarkPartOneWithUniquePathArray(b *testing.B) {
	input := util.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		partOneWithUniquePathArray(input)
	}
}

func BenchmarkPartTwoWithUniquePathArray(b *testing.B) {
	input := util.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		partTwoWithUniquePathArray(input)
	}
}

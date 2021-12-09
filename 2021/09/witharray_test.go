package main

import (
	"testing"

	"github.com/jdrst/adventofgo/util"
)

func TestPartOneWithArrays(t *testing.T) {
	expected := 15
	actual := partOneWithArrays(util.File(testInput).WithOSLinebreaks())
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}

	expected = 588
	actual = partOneWithArrays(util.ReadFile("input.txt"))
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

func TestPartTwoWithArrays(t *testing.T) {
	expected := 1134
	actual := partTwoWithArrays(util.File(testInput).WithOSLinebreaks())
	if actual != expected {
		t.Errorf("\nexpected was: %v\nactual is: %v", expected, actual)
	}

	expected = 964712
	actual = partTwoWithArrays(util.ReadFile("input.txt"))
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

func BenchmarkPartOneWithArrays(b *testing.B) {
	input := util.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		partOneWithArrays(input)
	}
}

func BenchmarkPartTwoWithArrays(b *testing.B) {
	input := util.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		partTwoWithArrays(input)
	}
}

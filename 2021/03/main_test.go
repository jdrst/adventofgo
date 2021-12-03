package main

import (
	"testing"

	"github.com/jdrst/adventofgo/util"
)

var testInput = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

// func TestPartOne(t *testing.T) {
// 	expected := "testOne"
// 	actual := partOne(util.File(testInput).WithOSLinebreaks())
// 	if actual != expected {
// 		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
// 	}
// }

func TestPartTwo(t *testing.T) {
	expected := 230
	actual := partTwo(util.File(testInput).WithOSLinebreaks())
	if actual != expected {
		t.Errorf("\nexpected1 was: %v\nactual1 is: %v", expected, actual)
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

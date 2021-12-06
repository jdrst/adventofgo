package main

import (
	"testing"

	"github.com/jdrst/adventofgo/util"
)

var testInput = `3,4,3,1,2`

func TestPartOne(t *testing.T) {
	expected := 26
	actual := partOne(util.File(testInput).WithOSLinebreaks(), 18)
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}

	expected = 5934
	actual = partOne(util.File(testInput).WithOSLinebreaks(), 80)
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

// func TestPartTwo(t *testing.T) {
// 	expected := "testTwo"
// 	actual := partTwo(util.File(testInput).WithOSLinebreaks())
// 	if actual != expected {
// 		t.Errorf("\nexpected was: %v\nactual is: %v", expected, actual)
// 	}
// }

func BenchmarkPartOne(b *testing.B) {
	input := util.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		partOne(input, 80)
	}
}

// func BenchmarkPartTwo(b *testing.B) {
// 	input := util.ReadFile("input.txt")
// 	for n := 0; n < b.N; n++ {
// 		partTwo(input)
// 	}
// }

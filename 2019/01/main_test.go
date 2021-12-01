package main

import (
	"testing"

	"git.threeman.info/jd/adventofcode/util"
)

var testInput = `12
14
1969
100756`

var testInput2 = `14
1969
100756`

func TestPartOne(t *testing.T) {
	expected := 2 + 2 + 654 + 33583
	actual := partOne(util.File(testInput).WithOSLinebreaks())
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 2 + 966 + 50346
	actual := partTwo(util.File(testInput2).WithOSLinebreaks())
	if actual != expected {
		t.Errorf("\nexpected was: %v\nactual is: %v", expected, actual)
	}
}

func BenchmarkMain(b *testing.B) {
	for n := 0; n < b.N; n++ {
		main()
	}
}

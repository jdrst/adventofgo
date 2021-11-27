package main

import (
	"testing"

	"git.threeman.info/jd/adventofcode/util"
)

var testInput = `testOne
testTwo`

func TestPartOne(t *testing.T) {
	expected := "testOne"
	actual := partOne(util.File(testInput).WithCRLF())
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := "testTwo"
	actual := partTwo(util.File(testInput).WithCRLF())
	if actual != expected {
		t.Errorf("\nexpected was: %v\nactual is: %v", expected, actual)
	}
}

func BenchmarkMain(b *testing.B) {
	for n := 0; n < b.N; n++ {
		main()
	}
}

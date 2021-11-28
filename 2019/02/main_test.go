package main

import (
	"testing"

	"git.threeman.info/jd/adventofcode/util"
)

var tests = []struct {
	input    string
	expected int
}{
	{`1,9,10,3,2,3,11,0,99,30,40,50`, 3500},
	{`1,0,0,0,99`, 2},
	{`2,3,0,3,99`, 2},
	{`2,4,4,5,99,0`, 2},
	{`1,1,1,4,99,5,6,0,99`, 30},
}

func TestPartOne(t *testing.T) {
	for _, test := range tests {
		expected := test.expected
		actual := partOne(util.File(test.input).WithCRLF())
		if actual != expected {
			t.Errorf("expected was: %v \n actual is: %v", expected, actual)
		}
	}
}

// func TestPartTwo(t *testing.T) {
// 	expected := "testTwo"
// 	actual := partTwo(util.File(testInput).WithCRLF())
// 	if actual != expected {
// 		t.Errorf("\nexpected was: %v\nactual is: %v", expected, actual)
// 	}
// }

func BenchmarkMain(b *testing.B) {
	for n := 0; n < b.N; n++ {
		main()
	}
}

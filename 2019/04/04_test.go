package main

import (
	"testing"
)

var testsPart1 = []struct {
	input    string
	expected bool
}{
	{`122345`, true},
	{`111111`, true},
	{`223450`, false},
	{`123789`, false},
}

var testsPart2 = []struct {
	input    string
	expected bool
}{
	{`112233`, true},
	{`111122`, true},
	{`111233`, true},
	{`113334`, true},
	{`344666`, true},
	{`133334`, false},
	{`111322`, false},
	{`123444`, false},
	{`124444`, false},
	{`144444`, false},
	{`444444`, false},
	{`113544`, false},
}

func TestPartOne(t *testing.T) {
	for _, test := range testsPart1 {
		actual := !hasDecreasingDigits([]rune(test.input)) && hasSameDigitAdjacent([]rune(test.input))
		if actual != test.expected {
			t.Errorf("input: %v\nexpected was: %v \nactual is: %v", test.input, test.expected, actual)
		}
	}
}

func TestPartTwo(t *testing.T) {
	for _, test := range testsPart2 {
		actual := !hasDecreasingDigits([]rune(test.input)) && hasSameDigitAdjacentWithoutLargerGroup([]rune(test.input))
		if actual != test.expected {
			t.Errorf("input: %v\nexpected was: %v \nactual is: %v", test.input, test.expected, actual)
		}
	}
}

func BenchmarkMain(b *testing.B) {
	for n := 0; n < b.N; n++ {
		main()
	}
}

package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestNumberN(t *testing.T) {
	for _, test := range testcases {
		lastSpoken, _ := strconv.Atoi(strings.Split(test.input, ",")[2])
		input := prepInput([]byte(test.input))
		if actual := numberN(input, lastSpoken, test.rounds); actual != test.expected {
			t.Errorf("numberN(%v, %d) = %d, expected %d.",
				test.input, lastSpoken, actual, test.expected)
		}
	}
}

func BenchmarkNumberN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testcases {
			lastSpoken, _ := strconv.Atoi(strings.Split(test.input, ",")[2])
			input := prepInput([]byte(test.input))
			numberN(input, lastSpoken, test.rounds)
		}
	}
}

var testcases = []struct {
	input    string
	rounds   int
	expected int
}{
	{"0,3,6", 2020, 436},
	{"1,3,2", 2020, 1},
	{"2,1,3", 2020, 10},
	{"1,2,3", 2020, 27},
	{"2,3,1", 2020, 78},
	{"3,2,1", 2020, 438},
	{"3,1,2", 2020, 1836},
	{"0,3,6", 30000000, 175594},
	{"1,3,2", 30000000, 2578},
	{"2,1,3", 30000000, 3544142},
	{"1,2,3", 30000000, 261214},
	{"2,3,1", 30000000, 6895259},
	{"3,2,1", 30000000, 18},
	{"3,1,2", 30000000, 362},
}

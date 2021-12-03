package main

import (
	"testing"
)

func TestNumberN(t *testing.T) {
	newLine = "\n"
	// for _, test := range testcases {
	// 	lastSpoken, _ := strconv.Atoi(strings.Split(test.input, ",")[2])
	// 	input := prepInput([]byte(test.input))
	// 	if actual := numberN(input, lastSpoken, test.rounds); actual != test.expected {
	// 		t.Errorf("numberN(%v, %d) = %d, expected %d.",
	// 			test.input, lastSpoken, actual, test.expected)
	// 	}
	// }
}

func BenchmarkNumberN(b *testing.B) {
	newLine = "\n"
	for i := 0; i < b.N; i++ {
		// for _, test := range testcases {
		// 	lastSpoken, _ := strconv.Atoi(strings.Split(test.input, ",")[2])
		// 	input := prepInput([]byte(test.input))
		// 	numberN(input, lastSpoken, test.rounds)
		// }
	}
}

var testcases = []struct {
	input    string
	rounds   int
	expected int
}{
	{"", 2020, 436},
}

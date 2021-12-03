package main

import (
	"testing"
)

func TestMultipliedDiffs(t *testing.T) {
	newLine = "\n"
	diffs := multipliedDiffs(prepInput([]byte(testinput1)))
	expected := 35
	if diffs != expected {
		t.Errorf("Product was incorrect, got: %d, want: %d.", diffs, expected)
	}
	diffs = multipliedDiffs(prepInput([]byte(testinput2)))
	expected = 220
	if diffs != expected {
		t.Errorf("Product was incorrect, got: %d, want: %d.", diffs, expected)
	}
}

func TestPossiblePaths(t *testing.T) {
	newLine = "\n"
	paths := possiblePaths(prepInput([]byte(testinput1)))
	expected := 8
	if paths != expected {
		t.Errorf("Number of possible paths was incorrect, got: %d, want: %d.", paths, expected)
	}
	paths = possiblePaths(prepInput([]byte(testinput2)))
	expected = 19208
	if paths != expected {
		t.Errorf("Number of possible paths was incorrect, got: %d, want: %d.", paths, expected)
	}
}

func BenchmarkPossiblePaths(b *testing.B) {
	newLine = "\r\n"
	ints := prepInput(parseInput())
	for n := 0; n < b.N; n++ {
		possiblePaths(ints)
	}
}
func BenchmarkMultipliedDiffs(b *testing.B) {
	newLine = "\r\n"
	ints := prepInput(parseInput())
	for n := 0; n < b.N; n++ {
		multipliedDiffs(ints)
	}
}

const testinput1 = `16
10
15
5
1
11
7
19
6
12
4`

const testinput2 = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

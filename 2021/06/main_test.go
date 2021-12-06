package main

import (
	"testing"

	"github.com/jdrst/adventofgo/util"
)

var testInput = `3,4,3,1,2`

func TestLanternfishSpawnedAfter(t *testing.T) {
	expected := 26
	actual := lanternfishSpawnedAfter(18, util.File(testInput).WithOSLinebreaks())
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}

	expected = 5934
	actual = lanternfishSpawnedAfter(80, util.File(testInput).WithOSLinebreaks())
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

func BenchmarkLanternfishSpawnedAfter(b *testing.B) {
	input := util.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		lanternfishSpawnedAfter(256, input)
	}
}

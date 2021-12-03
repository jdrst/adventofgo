package main

import (
	"math/big"
	"strings"
	"testing"
	"github.com/deanveloper/modmath/v1/bigmod"
)

// func TestPart1(t *testing.T) {
// 	newLine = "\n"
// 	actual := 0
// 	expected := 35
// 	if actual != expected {
// 		t.Errorf("blalba, got: %d, want: %d.", actual, expected)
// 	}
// 	actual = 0
// 	expected = 220
// 	if actual != expected {
// 		t.Errorf("blalba, got: %d, want: %d.", actual, expected)
// 	}
// }

func TestPart2(t *testing.T) {
	offsets1 := makeOffsets(strings.Split(testinput1, ","))
	offsets2 := makeOffsets(strings.Split(testinput2, ","))
	offsets3 := makeOffsets(strings.Split(testinput3, ","))
	offsets4 := makeOffsets(strings.Split(testinput4, ","))
	offsets5 := makeOffsets(strings.Split(testinput5, ","))
	actual := bigmod.SolveCrtMany(makeCrtMat(offsets1))
	expected := big.NewInt(3417)
	if actual.Cmp(expected) != 0 {
		t.Errorf("Expected the earliest timestamp to be %d, got: %d.", expected, actual)
	}
	actual = bigmod.SolveCrtMany(makeCrtMat(offsets2))
	expected = big.NewInt(754018)
	if actual.Cmp(expected) != 0 {
		t.Errorf("Expected the earliest timestamp to be %d, got: %d.", expected, actual)
	}
	actual = bigmod.SolveCrtMany(makeCrtMat(offsets3))
	expected = big.NewInt(779210)
	if actual.Cmp(expected) != 0 {
		t.Errorf("Expected the earliest timestamp to be %d, got: %d.", expected, actual)
	}
	actual = bigmod.SolveCrtMany(makeCrtMat(offsets4))
	expected = big.NewInt(1261476)
	if actual.Cmp(expected) != 0 {
		t.Errorf("Expected the earliest timestamp to be %d, got: %d.", expected, actual)
	}
	actual = bigmod.SolveCrtMany(makeCrtMat(offsets5))
	expected = big.NewInt(1202161486)
	if actual.Cmp(expected) != 0 {
		t.Errorf("Expected the earliest timestamp to be %d, got: %d.", expected, actual)
	}
}

func BenchmarkMain(b *testing.B) {
	newLine = "\r\n"
	//input := prepInput(parseInput())
	for n := 0; n < b.N; n++ {
		main()
	}
}

const testinput1 = `17,x,13,19`
const testinput2 = `67,7,59,61`
const testinput3 = `67,x,7,59,61`
const testinput4 = `67,7,x,59,61`
const testinput5 = `1789,37,47,1889`

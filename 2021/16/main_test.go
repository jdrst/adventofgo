package main

import (
	"testing"

	"github.com/jdrst/adventofgo/util"
)

func TestPartOne(t *testing.T) {
	var testCases = []struct {
		expected int
		input    string
	}{
		{16, `8A004A801A8002F478`},
		{12, `620080001611562C8802118E34`},
		{23, `C0015000016115A2E0802F182340`},
		{31, `A0016C880162017C3686B18A3D4780`},
	}
	for _, test := range testCases {
		actual := partOne(util.File(test.input).WithOSLinebreaks())
		if actual != test.expected {
			t.Errorf("\nexpected was: %v\nactual is: %v", test.expected, actual)
		}
	}

	expected := 960
	actual := partOne(util.ReadFile("input.txt"))
	if actual != expected {
		t.Errorf("\nexpected was: %v\nactual is: %v", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	var testCases = []struct {
		expected int
		input    string
	}{
		{3, `C200B40A82`},
		{54, `04005AC33890`},
		{7, `880086C3E88112`},
		{9, `CE00C43D881120`},
		{1, `D8005AC2A8F0`},
		{0, `F600BC2D8F`},
		{0, `9C005AC2F8F0`},
		{1, `9C0141080250320F1802104A08`},
		{1, `38006F45291200`},
		{3, `EE00D40C823060`},
		{2021, `D2FE28`},
	}
	for _, test := range testCases {
		actual := partTwo(util.File(test.input).WithOSLinebreaks())
		if actual != test.expected {
			t.Errorf("\nexpected was: %v\nactual is: %v", test.expected, actual)
		}
	}

	expected := 12301926782560
	actual := partTwo(util.ReadFile("input.txt"))
	if actual != expected {
		t.Errorf("\nexpected was: %v\nactual is: %v", expected, actual)
	}
}

func BenchmarkPartOne(b *testing.B) {
	input := util.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		partOne(input)
	}
}

func BenchmarkToBinaryString(b *testing.B) {
	input := util.ReadFile("input.txt").AsLines()
	for n := 0; n < b.N; n++ {
		toBinaryString(string(input[0]))
	}
}
func BenchmarkPartTwo(b *testing.B) {
	input := util.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		partTwo(input)
	}
}

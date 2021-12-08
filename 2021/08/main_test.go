package main

import (
	"testing"

	"github.com/jdrst/adventofgo/util"
)

var testInput = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

func TestPartOne(t *testing.T) {
	expected := 26
	actual := partOne(util.File(testInput).WithOSLinebreaks())
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 61229
	actual := partTwo(util.File(testInput).WithOSLinebreaks())
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

func BenchmarkPartTwo(b *testing.B) {
	input := util.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		partTwo(input)
	}
}
func TestDeduce(t *testing.T) {
	pattern := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab"
	output := "cdfeb fcadb cdfeb cdbaf"
	expected := 5353
	actual := deduceValue(pattern, output)
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

var IsInAnyOrderTestcases = []struct {
	first    string
	second   string
	expected bool
}{
	{first: "aefg", second: "gafe", expected: true},
	{first: "aefg", second: "aefd", expected: false},
	{first: "aefg", second: "aefgg", expected: false},
}

func TestIsInAnyOrder(t *testing.T) {
	for _, test := range IsInAnyOrderTestcases {
		if IsInAnyOrder(test.first, test.second) != test.expected {
			t.Errorf("IsInAnyOrder for %v and %v should return %v", test.first, test.second, test.expected)
		}
	}
}

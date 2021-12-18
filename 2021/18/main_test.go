package main

import (
	"fmt"
	"testing"

	"github.com/jdrst/adventofgo/util"
)

var testInput = `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`

func TestPartOne(t *testing.T) {
	expected := 4140
	actual := partOne(util.File(testInput).WithOSLinebreaks())
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

// func TestSplitRecursive(t *testing.T) {
// 	in := "[[[[4,0],[5,4]],[[7,7],[6,0]]],[17,[[11,9],[11,0]]]]"
// 	num := toSfnum(in)
// 	num.splitRecursive()
// }

func TestReduceRecursive2(t *testing.T) {
	var testCases = []struct {
		input    string
		expected string
	}{
		{"[[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]", "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]"},
	}
	for _, test := range testCases {
		num := toSfnum(test.input)
		num.reduce()
		actual := num.String()
		if test.expected != actual {
			t.Errorf("\nexpected was: %v\nactual is: %v", test.expected, actual)
		}
	}
}

func TestPartTwo(t *testing.T) {
	expected := 3993
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

func TestReduceRecursive(t *testing.T) {
	var testCases = []struct {
		input    string
		expected int
	}{
		// explode only
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", 633},
		{"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", 1384},
		{"[[[[[9,8],1],2],3],4]", 548},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", 633},
		{"[7,[6,[5,[4,[3,2]]]]]", 285},
		{"[[6,[5,[4,[3,2]]]],1]", 402},
		// {"[[6,[[[6,2],[5,6]],[[7,6],[4,7]]]],[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]]", }
		//
		/*
			[[[[0,7],4],[[7,8],[6,0]]],[8,1]]
			[[[14,4],[37,18]],26]
			[[50,147],26]
			[444,26]
			1384
		*/
	}
	for _, test := range testCases {
		num := toSfnum(test.input)
		num.reduce()
		// num.reduce()
		actual := *num.magnitude()
		if test.expected != actual {
			t.Errorf("\nexpected was: %v\nactual is: %v", test.expected, actual)
		}
	}
}

func TestStringer(t *testing.T) {
	var testCases = []struct {
		input string
	}{
		// explode only
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"},
		{"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"},
		{"[[[[[9,8],1],2],3],4]"},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"[7,[6,[5,[4,[3,2]]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]"},
		// {"[[6,[[[6,2],[5,6]],[[7,6],[4,7]]]],[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]]", }
		//
		/*
			[[[[0,7],4],[[7,8],[6,0]]],[8,1]]
			[[[14,4],[37,18]],26]
			[[50,147],26]
			[444,26]
			1384
		*/
	}
	for _, test := range testCases {
		num := toSfnum(test.input)
		actual := fmt.Sprint(num)
		if test.input != actual {
			t.Errorf("\nexpected was: %v\nactual is: %v", test.input, actual)
		}
	}
}

// func TestReduce(t *testing.T) {
// 	var testCases = []struct {
// 		input    string
// 		expected int
// 	}{
// 		// explode only
// 		// {"[[[[[9,8],1],2],3],4]", 548},
// 		{"[7,[6,[5,[4,[3,2]]]]]", 285},
// 		// {"[[6,[5,[4,[3,2]]]],1]", 602},
// 		// {"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", 633},
// 		//
// 		/*
// 			[[[[0,7],4],[[7,8],[6,0]]],[8,1]]
// 			[[[14,4],[37,18]],26]
// 			[[50,147],26]
// 			[444,26]
// 			1384
// 		*/
// 		// {"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", 1384},
// 	}
// 	for _, test := range testCases {
// 		num := toSfnum(test.input)
// 		num.reduce()
// 		actual := *num.magnitude()
// 		if test.expected != actual {
// 			t.Errorf("\nexpected was: %v\nactual is: %v", test.expected, actual)
// 		}
// 	}
// }

func TestCalcMagnitude(t *testing.T) {
	var testCases = []struct {
		input    string
		expected int
	}{
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
	}
	for _, test := range testCases {
		actual := toSfnum(test.input).magnitude()
		if test.expected != *actual {
			t.Errorf("\nexpected was: %v\nactual is: %v", test.expected, actual)
		}
	}
}

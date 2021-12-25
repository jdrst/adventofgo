package main

import (
	"strings"
	"testing"

	"github.com/jdrst/adventofgo/util"
)

var testInput = `inp x
mul x -1
add z x`

var testInput2 = `inp z
inp x
mul z 3
eql z x`

var testInput3 = `inp w
add z w
mod z 2
div w 2
add y w
mod y 2
div w 2
add x w
mod x 2
div w 2
mod w 2
mul y 10
add z y
mul x 100
add z x
mul w 1000
add z w`

func TestALUOne(t *testing.T) {
	expected := -1
	lines := strings.Split(string(util.File(testInput).WithOSLinebreaks()), util.NewLine())

	var w, x, y, z int
	p := ALU{instructions: lines, w: &w, x: &x, y: &y, z: &z}

	actual := p.process(1)
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}

	expected = -9
	actual = p.process(9)
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

func TestALUTwo(t *testing.T) {
	expected := 1
	lines := strings.Split(string(util.File(testInput2).WithOSLinebreaks()), util.NewLine())

	var w, x, y, z int
	p := ALU{instructions: lines, w: &w, x: &x, y: &y, z: &z}

	actual := p.process(39)
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}

	expected = 0
	actual = p.process(98)
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

func TestALUThree(t *testing.T) {
	expected := 1000
	lines := strings.Split(string(util.File(testInput3).WithOSLinebreaks()), util.NewLine())

	var w, x, y, z int
	p := ALU{instructions: lines, w: &w, x: &x, y: &y, z: &z}

	actual := p.process(8)
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}

	expected = 101
	actual = p.process(5)
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

func TestPartOne(t *testing.T) {
	expected := 51939397989999

	actual := partOne(util.ReadFile("input.txt"))
	if actual != expected {
		t.Errorf("expected was: %v \n actual is: %v", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 11717131211195
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

func BenchmarkPartTwo(b *testing.B) {
	input := util.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		partTwo(input)
	}
}

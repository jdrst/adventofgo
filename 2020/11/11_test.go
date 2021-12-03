package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	newLine = "\n"
	tiles := prepareInput([]byte(startingMap))
	floor := makeSeats(tiles)
	actual := floor.occupied()
	expected := 0
	if actual != expected {
		t.Errorf("Expected initial seatmap to have no occupied seats, got: %d, want: %d.", actual, expected)
	}
	changeSeats(&floor, rule01)
	actual = floor.occupied()
	expected = strings.Count(p1round1, "#")
	if actual != expected {
		t.Errorf("Expected first seat changing to have all seats occupied, got: %d, want: %d.", actual, expected)
	}
	changeSeats(&floor, rule01)
	actual = floor.occupied()
	expected = strings.Count(p1round2, "#")
	if actual != expected {
		t.Errorf("Expected second seat changing to have this number of seats occupied: got: %d, want: %d.", actual, expected)
	}
	changeSeats(&floor, rule01)
	actual = floor.occupied()
	expected = strings.Count(p1round3, "#")
	if actual != expected {
		t.Errorf("Expected third seat changing to have this number of seats occupied: got: %d, want: %d.", actual, expected)
	}
	changeSeats(&floor, rule01)
	actual = floor.occupied()
	expected = strings.Count(p1round4, "#")
	if actual != expected {
		t.Errorf("Expected fourth seat changing to have this number of seats occupied: got: %d, want: %d.", actual, expected)
	}
	changeSeats(&floor, rule01)
	actual = floor.occupied()
	expected = strings.Count(p1round5, "#")
	if actual != expected {
		t.Errorf("Expected fifth seat changing to have this number of seats occupied: got: %d, want: %d.", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	newLine = "\n"
	tiles := prepareInput([]byte(startingMap))
	floor := makeSeats(tiles)
	actual := floor.occupied()
	expected := 0
	if actual != expected {
		t.Errorf("Expected initial seatmap to have no occupied seats, got: %d, want: %d.", actual, expected)
	}
	changeSeats(&floor, rule02)
	actual = floor.occupied()
	expected = strings.Count(p2round1, "#")
	if actual != expected {
		t.Errorf("Expected first seat changing to have all seats occupied, got: %d, want: %d.", actual, expected)
	}
	changeSeats(&floor, rule02)
	actual = floor.occupied()
	expected = strings.Count(p2round2, "#")
	if actual != expected {
		t.Errorf("Expected second seat changing to have this number of seats occupied: got: %d, want: %d.", actual, expected)
	}
	changeSeats(&floor, rule02)
	actual = floor.occupied()
	expected = strings.Count(p2round3, "#")
	if actual != expected {
		t.Errorf("Expected third seat changing to have this number of seats occupied: got: %d, want: %d.", actual, expected)
	}
	changeSeats(&floor, rule02)
	actual = floor.occupied()
	expected = strings.Count(p2round4, "#")
	if actual != expected {
		t.Errorf("Expected fourth seat changing to have this number of seats occupied: got: %d, want: %d.", actual, expected)
	}
	changeSeats(&floor, rule02)
	actual = floor.occupied()
	expected = strings.Count(p2round5, "#")
	if actual != expected {
		t.Errorf("Expected fifth seat changing to have this number of seats occupied: got: %d, want: %d.", actual, expected)
	}
	changeSeats(&floor, rule02)
	actual = floor.occupied()
	expected = strings.Count(p2round6, "#")
	if actual != expected {
		t.Errorf("Expected sixth seat changing to have this number of seats occupied: got: %d, want: %d.", actual, expected)
	}
}

func BenchmarkMain(b *testing.B) {
	newLine = "\r\n"
	for n := 0; n < b.N; n++ {
		main()
	}
}

const startingMap = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

const p1round1 = `#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`

const p1round2 = `#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##`

const p1round3 = `#.##.L#.##
#L###LL.L#
L.#.#..#..
#L##.##.L#
#.##.LL.LL
#.###L#.##
..#.#.....
#L######L#
#.LL###L.L
#.#L###.##`

const p1round4 = `#.#L.L#.##
#LLL#LL.L#
L.L.L..#..
#LLL.##.L#
#.LL.LL.LL
#.LL#L#.##
..L.L.....
#L#LLLL#L#
#.LLLLLL.L
#.#L#L#.##`

const p1round5 = `#.#L.L#.##
#LLL#LL.L#
L.#.L..#..
#L##.##.L#
#.#L.LL.LL
#.#L#L#.##
..L.L.....
#L#L##L#L#
#.LLLLLL.L
#.#L#L#.##`

const p2round1 = `#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`

const p2round2 = `#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#`

const p2round3 = `#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#`

const p2round4 = `#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##LL.LL.L#
L.LL.LL.L#
#.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLL#.L
#.L#LL#.L#`

const p2round5 = `#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.#L.L#
#.L####.LL
..#.#.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`

const p2round6 = `#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`

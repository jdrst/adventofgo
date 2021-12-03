package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type position struct {
	x, y int
}

type seatLayout struct {
	rowLength, colLength int
	grid                 map[position]bool
}

type seatingRule struct {
	stayWhile             int
	occupiedSeatsSeenFrom occupiedSeatsSeenFunc
}

type occupiedSeatsSeenFunc = func(position, *seatLayout) int

var rule01 = seatingRule{3, rule01Func}
var rule02 = seatingRule{4, rule02Func}

var newLine = "\r\n"

var directions = []position{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func main() {
	tiles := prepareInput(parseInput())

	seats := makeSeats(tiles)

	for c := 1; c > 0; {
		c = changeSeats(&seats, rule01)
	}
	fmt.Println("01")
	fmt.Println(seats.occupied())

	seats = makeSeats(tiles)

	for c := 1; c > 0; {
		c = changeSeats(&seats, rule02)
	}
	fmt.Println("02")
	fmt.Println(seats.occupied())
}

func changeSeats(seats *seatLayout, rule seatingRule) int {
	newGrid := make(map[position]bool)
	changes := 0
	for pos, occ := range seats.grid {
		switch occs := rule.occupiedSeatsSeenFrom(pos, seats); {
		case occs == 0 && !occ:
			changes++
			newGrid[pos] = true
		case occs > rule.stayWhile && occ:
			changes++
			newGrid[pos] = false
		default:
			newGrid[pos] = occ
		}
	}
	seats.grid = newGrid
	return changes
}

func rule02Func(pos position, seats *seatLayout) int {
	occs := 0

directions:
	for _, ver := range directions {
		for i, j := pos.x+ver.x, pos.y+ver.y; i < seats.rowLength && i >= 0 && j < seats.colLength && j >= 0; i, j = i+ver.x, j+ver.y {
			if occ, ok := seats.grid[position{i, j}]; ok {
				if occ {
					occs++
				}
				continue directions
			}
		}
	}
	return occs
}

func rule01Func(pos position, seats *seatLayout) int {
	occs := 0
	for _, ver := range directions {
		if occ, ok := seats.grid[position{pos.x + ver.x, pos.y + ver.y}]; ok {
			if occ {
				occs++
			}
		}
	}
	return occs
}

func makeSeats(tiles [][]string) seatLayout {
	seatmap := make(map[position]bool)
	rowLen := len(tiles)
	colLen := len(tiles[0])
	for i, row := range tiles {
		for j, tile := range row {
			if isSeat(tile) {
				seatmap[position{i, j}] = false
			}
		}
	}
	return seatLayout{rowLen, colLen, seatmap}
}

func isSeat(tile string) bool {
	return tile == "L" || tile == "#"
}

func (seats *seatLayout) occupied() int {
	count := 0
	for _, s := range seats.grid {
		if s {
			count++
		}
	}
	return count
}

func parseInput() []byte {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return input
}

func prepareInput(input []byte) [][]string {
	lines := strings.Split(strings.TrimSpace(string(input)), newLine)
	tiles := make([][]string, len(lines))

	for i, l := range lines {
		tiles[i] = strings.Split(l, "")
	}

	return tiles
}

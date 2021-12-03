package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}
type ship struct {
	position  point
	direction point
}

func main() {
	instructions := prepareInput(parseInput())
	ferry := ship{point{0, 0}, point{1, 0}}

	for _, instr := range instructions {
		ferry.move(instr, false)
	}
	fmt.Println(abs(ferry.position.x) + abs(ferry.position.y))

	ferry = ship{point{0, 0}, point{10, 1}}

	for _, instr := range instructions {
		ferry.move(instr, true)
	}
	fmt.Println(abs(ferry.position.x) + abs(ferry.position.y))
}

func (pos *ship) move(instruction string, relative bool) {

	to := instruction[0:1]
	amount, err := strconv.Atoi(instruction[1:])
	if err != nil {
		panic(err)
	}
	movingPos := &pos.position
	if relative {
		movingPos = &pos.direction
	}
	switch to {
	case "F":
		pos.position.x += pos.direction.x * amount
		pos.position.y += pos.direction.y * amount
	case "L":
		for amount > 0 {
			amount -= 90
			pos.direction.x, pos.direction.y = -pos.direction.y, pos.direction.x
		}
	case "R":
		for amount > 0 {
			amount -= 90
			pos.direction.x, pos.direction.y = pos.direction.y, -pos.direction.x
		}
	case "N":
		movingPos.y += amount
	case "E":
		movingPos.x += amount
	case "S":
		movingPos.y -= amount
	case "W":
		movingPos.x -= amount
	}
}

func parseInput() []byte {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return input
}

var newLine = "\r\n"

func prepareInput(input []byte) []string {
	return strings.Split(strings.TrimSpace(string(input)), newLine)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

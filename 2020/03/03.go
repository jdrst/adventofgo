package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func parseInput() []string {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(input)), "\r\n")
}

type vector struct {
	x, y int
}

func main() {
	grid := parseInput()
	fmt.Println(ride(grid, vector{x: 3, y: 1}))
	fmt.Println(ride(grid, vector{x: 1, y: 1}) * ride(grid, vector{x: 3, y: 1}) * ride(grid, vector{x: 5, y: 1}) * ride(grid, vector{x: 7, y: 1}) * ride(grid, vector{x: 1, y: 2}))
}

func ride(grid []string, trajectory vector) int {
	width := len(grid[0])
	x := 0
	y := 0
	trees := 0
	for y < len(grid) {
		line := strings.Split(grid[y], "")
		x = x % width
		if line[x] == "#" {
			trees++
		}
		x = x + trajectory.x
		y = y + trajectory.y
	}
	return trees
}

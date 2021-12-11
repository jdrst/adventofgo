package main

import (
	"fmt"

	"github.com/jdrst/adventofgo/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

type point struct {
	x, y int
}

type octopus struct {
	energy  int
	flashed bool
}

func partOne(file util.File) int {
	lines := file.AsLines()
	octopuses := [10][10]*octopus{}
	for i, l := range lines {
		for j, o := range l.SubSplitWith("").AsInts() {
			octopuses[i][j] = &octopus{energy: o, flashed: false}
		}
	}
	flashes := 0
	for i := 0; i < 100; i++ {
		flashes += doStep(octopuses)
	}
	return flashes
}

func partTwo(file util.File) int {
	lines := file.AsLines()
	octopuses := [10][10]*octopus{}
	for i, l := range lines {
		for j, o := range l.SubSplitWith("").AsInts() {
			octopuses[i][j] = &octopus{energy: o, flashed: false}
		}
	}
	i := 0
	for {
		for _, l := range octopuses {
			for _, o := range l {
				if !o.flashed {
					goto cont
				}
			}
		}
		return i
	cont:
		doStep(octopuses)
		i++
	}
}

func doStep(octopuses [10][10]*octopus) int {
	flashes := 0
	for _, l := range octopuses {
		for _, o := range l {
			if o.flashed {
				o.energy = 0
			}
			o.energy++
			o.flashed = false
		}
	}
	hasFlashed := true
	for hasFlashed {
		hasFlashed = false
		for i, l := range octopuses {
			for j, o := range l {
				if o.energy > 9 && !o.flashed {
					flashes++
					hasFlashed = true
					o.flashed = true
					for _, p := range adjacent(i, j) {
						octopuses[p.x][p.y].energy = octopuses[p.x][p.y].energy + 1
					}
				}
			}
		}
	}
	return flashes
}

func adjacent(x, y int) []point {
	res := make([]point, 0)
	if x > 0 {
		res = append(res, point{x - 1, y})
		if y > 0 {
			res = append(res, point{x - 1, y - 1})
		}
	}
	if x < 9 {
		res = append(res, point{x + 1, y})
		if y < 9 {
			res = append(res, point{x + 1, y + 1})
		}
	}
	if y < 9 {
		res = append(res, point{x, y + 1})
		if x > 0 {
			res = append(res, point{x - 1, y + 1})
		}
	}
	if y > 0 {
		res = append(res, point{x, y - 1})
		if x < 9 {
			res = append(res, point{x + 1, y - 1})
		}
	}
	return res
}

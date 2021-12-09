package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/jdrst/adventofgo/util"
)

type point struct {
	x, y int
}

type heightMap map[point]int

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	lines := file.AsLines()
	heightmap := make(heightMap)
	for i, l := range lines {
		for j, v := range l.SubSplitWith("").AsInts() {
			heightmap[point{i, j}] = v
		}
	}

	risklevel := 0

	for p := range heightmap {
		if heightmap.isLowPoint(p) {
			risklevel += heightmap[p] + 1
		}
	}

	return risklevel
}

func partTwo(file util.File) int {
	lines := file.AsLines()
	heightmap := make(heightMap)
	for i, l := range lines {
		for j, v := range l.SubSplitWith("").AsInts() {
			heightmap[point{i, j}] = v
		}
	}

	basins := make([]int, 0)

	for p := range heightmap {
		if heightmap.isLowPoint(p) {
			basins = append(basins, heightmap.basinSizeFor(p, make(map[point]bool)))
		}
	}

	sort.Ints(basins)
	return basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
}

func (hm heightMap) basinSizeFor(p point, visited map[point]bool) int {
	if _, exists := visited[p]; exists || hm.isBasinEnd(p) {
		return 0
	}
	visited[p] = true
	upper, lower, left, right := p.up(), p.down(), p.left(), p.right()
	return 1 + hm.basinSizeFor(upper, visited) + hm.basinSizeFor(lower, visited) + hm.basinSizeFor(left, visited) + hm.basinSizeFor(right, visited)
}

func (hm heightMap) isLowPoint(p point) bool {
	val := hm[p]
	upper := hm.valueOrMaxInt(p.up())
	lower := hm.valueOrMaxInt(p.down())
	left := hm.valueOrMaxInt(p.left())
	right := hm.valueOrMaxInt(p.right())
	return upper > val && left > val && lower > val && right > val
}

func (hm heightMap) valueOrMaxInt(p point) int {
	if v, exists := hm[p]; exists {
		return v
	}
	return math.MaxInt
}

func (heightmap heightMap) isBasinEnd(p point) bool {
	if v, exists := heightmap[p]; exists {
		return v == 9
	}
	return true
}

func (p point) up() point {
	return point{p.x + 1, p.y}
}

func (p point) down() point {
	return point{p.x - 1, p.y}
}

func (p point) left() point {
	return point{p.x, p.y - 1}
}

func (p point) right() point {
	return point{p.x, p.y + 1}
}

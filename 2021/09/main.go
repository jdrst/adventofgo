package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/jdrst/adventofgo/util"
)

type Point struct {
	x, y int
}

type Heightmap map[Point]int

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	lines := file.AsLines()
	heightmap := make(Heightmap)
	for i, l := range lines {
		for j, v := range l.SubSplitWith("").AsInts() {
			heightmap[Point{i, j}] = v
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
	heightmap := make(Heightmap)
	for i, l := range lines {
		for j, v := range l.SubSplitWith("").AsInts() {
			heightmap[Point{i, j}] = v
		}
	}

	basins := make([]int, 0)

	for p := range heightmap {
		if heightmap.isLowPoint(p) {
			basins = append(basins, heightmap.basinSizeFor(p, make(map[Point]bool)))
		}
	}

	sort.Ints(basins)
	return basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
}

func (hm Heightmap) basinSizeFor(p Point, visited map[Point]bool) int {
	if _, exists := visited[p]; exists {
		return 0
	}
	sum := 1
	visited[p] = true
	upper, lower, left, right := p.up(), p.down(), p.left(), p.right()
	if !hm.isBasinEnd(upper) {
		sum += hm.basinSizeFor(upper, visited)
	}
	if !hm.isBasinEnd(lower) {
		sum += hm.basinSizeFor(lower, visited)
	}
	if !hm.isBasinEnd(left) {
		sum += hm.basinSizeFor(left, visited)
	}
	if !hm.isBasinEnd(right) {
		sum += hm.basinSizeFor(right, visited)
	}
	return sum
}

func (hm Heightmap) isLowPoint(p Point) bool {
	val := hm[p]
	upper := hm.valueOrMaxInt(p.up())
	lower := hm.valueOrMaxInt(p.down())
	left := hm.valueOrMaxInt(p.left())
	right := hm.valueOrMaxInt(p.right())
	return upper > val && left > val && lower > val && right > val
}

func (hm Heightmap) valueOrMaxInt(p Point) int {
	if v, exists := hm[p]; exists {
		return v
	}
	return math.MaxInt
}

func (heightmap Heightmap) isBasinEnd(p Point) bool {
	if v, exists := heightmap[p]; exists {
		return v == 9
	}
	return true
}

func (p Point) up() Point {
	return Point{p.x + 1, p.y}
}

func (p Point) down() Point {
	return Point{p.x - 1, p.y}
}

func (p Point) left() Point {
	return Point{p.x, p.y - 1}
}

func (p Point) right() Point {
	return Point{p.x, p.y + 1}
}

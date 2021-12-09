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

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	lines := file.AsLines()
	heightmap := make(map[Point]int)
	for i, l := range lines {
		for j, v := range l.SubSplitWith("").AsInts() {
			heightmap[Point{i, j}] = v
		}
	}

	risklevel := 0

	for p := range heightmap {
		if isLowPoint(p, heightmap) {
			risklevel += heightmap[p] + 1
		}
	}

	return risklevel
}

func partTwo(file util.File) int {
	lines := file.AsLines()
	heightmap := make(map[Point]int)
	for i, l := range lines {
		for j, v := range l.SubSplitWith("").AsInts() {
			heightmap[Point{i, j}] = v
		}
	}

	basins := make([]int, 0)

	for p := range heightmap {
		if isLowPoint(p, heightmap) {
			basins = append(basins, getBasinSize(p, heightmap, make(map[Point]bool)))
		}
	}

	sort.Ints(basins)
	return basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
}

func isLowPoint(p Point, heightmap map[Point]int) bool {
	val := heightmap[p]
	upper, lower, left, right := math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt
	if v, exists := heightmap[Point{p.x + 1, p.y}]; exists {
		upper = v
	}
	if v, exists := heightmap[Point{p.x - 1, p.y}]; exists {
		lower = v
	}
	if v, exists := heightmap[Point{p.x, p.y + 1}]; exists {
		left = v
	}
	if v, exists := heightmap[Point{p.x, p.y - 1}]; exists {
		right = v
	}
	return upper > val && left > val && lower > val && right > val
}

func getBasinSize(p Point, heightmap map[Point]int, visited map[Point]bool) int {
	if _, exists := visited[p]; exists {
		return 0
	}
	sum := 1
	visited[p] = true
	upper, lower, left, right := Point{p.x + 1, p.y}, Point{p.x - 1, p.y}, Point{p.x, p.y + 1}, Point{p.x, p.y - 1}
	if !isBasinEnd(upper, heightmap) {
		sum += getBasinSize(upper, heightmap, visited)
	}
	if !isBasinEnd(lower, heightmap) {
		sum += getBasinSize(lower, heightmap, visited)
	}
	if !isBasinEnd(left, heightmap) {
		sum += getBasinSize(left, heightmap, visited)
	}
	if !isBasinEnd(right, heightmap) {
		sum += getBasinSize(right, heightmap, visited)
	}
	return sum
}

func isBasinEnd(p Point, heightmap map[Point]int) bool {
	if v, exists := heightmap[p]; exists {
		return v == 9
	}
	return true
}

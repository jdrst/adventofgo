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
	heightmap := make([][]int, len(lines))
	for i, l := range lines {
		heightmap[i] = l.SubSplitWith("").AsInts()
	}

	risklevel := 0

	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			if isLowPoint(i, j, heightmap) {
				risklevel += heightmap[i][j] + 1
			}
		}
	}

	return risklevel
}

func partTwo(file util.File) int {
	lines := file.AsLines()
	heightmap := make([][]int, len(lines))
	for i, l := range lines {
		heightmap[i] = l.SubSplitWith("").AsInts()
	}

	basins := make([]int, 0)

	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			if isLowPoint(i, j, heightmap) {
				basins = append(basins, getBasinSize(Point{i, j}, ToHashMap(heightmap), make(map[Point]bool)))
			}
		}
	}

	sort.Ints(basins)
	return basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
}

func isLowPoint(x, y int, heightmap [][]int) bool {
	val := heightmap[x][y]
	upper, lower, left, right := math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt
	if y > 0 {
		upper = heightmap[x][y-1]
	}
	if x > 0 {
		left = heightmap[x-1][y]
	}
	if x < len(heightmap)-1 {
		right = heightmap[x+1][y]
	}
	if y < len(heightmap[x])-1 {
		lower = heightmap[x][y+1]
	}
	return upper > val && left > val && lower > val && right > val
}

func ToHashMap(heightmap [][]int) map[Point]int {
	res := make(map[Point]int)
	for i, r := range heightmap {
		for j, v := range r {
			res[Point{i, j}] = v
		}
	}
	return res
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

// func isVisited(p Point, calculated map[Point]bool) bool {
// 	_, exists := calculated[p]
// 	return exists
// }

func isBasinEnd(p Point, heightmap map[Point]int) bool {
	if v, exists := heightmap[p]; exists {
		return v == 9
	}
	return true
}

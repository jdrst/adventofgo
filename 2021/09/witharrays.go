package main

import (
	"math"
	"sort"

	"github.com/jdrst/adventofgo/util"
)

type node struct {
	value   int
	visited bool
}

type heightArray [][]node

func partOneWithArrays(file util.File) int {
	lines := file.AsLines()
	heightArray := make(heightArray, len(lines))
	for i, l := range lines {
		heightArray[i] = make([]node, len(l))
		for j, v := range l.SubSplitWith("").AsInts() {
			heightArray[i][j] = node{value: v, visited: false}
		}
	}

	risklevel := 0

	for i, r := range heightArray {
		for j, n := range r {
			if heightArray.isLowPoint(i, j) {
				risklevel += 1 + n.value
			}
		}
	}

	return risklevel
}

func partTwoWithArrays(file util.File) int {
	lines := file.AsLines()
	heightArray := make(heightArray, len(lines))
	for i, l := range lines {
		heightArray[i] = make([]node, len(l))
		for j, v := range l.SubSplitWith("").AsInts() {
			heightArray[i][j] = node{value: v, visited: false}
		}
	}

	basins := make([]int, 0)

	for i, r := range heightArray {
		for j := range r {
			if heightArray.isLowPoint(i, j) {
				basins = append(basins, heightArray.basinSizeFor(i, j))
			}
		}
	}

	sort.Ints(basins)
	return basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
}

func (ha heightArray) basinSizeFor(x, y int) int {
	if ha.isOutOfBounds(x, y) || ha[x][y].visited {
		return 0
	}
	sum := 1
	ha[x][y].visited = true
	nx, ny := x+1, y
	if !ha.isBasinEnd(nx, ny) {
		sum += ha.basinSizeFor(nx, ny)
	}
	nx = x - 1
	if !ha.isBasinEnd(nx, ny) {
		sum += ha.basinSizeFor(nx, ny)
	}
	nx, ny = x, y+1
	if !ha.isBasinEnd(nx, ny) {
		sum += ha.basinSizeFor(nx, ny)
	}
	ny = y - 1
	if !ha.isBasinEnd(nx, ny) {
		sum += ha.basinSizeFor(nx, ny)
	}
	return sum
}

func (ha heightArray) isLowPoint(x, y int) bool {
	val := ha[x][y].value
	upper := ha.valueOrMaxInt(x+1, y)
	lower := ha.valueOrMaxInt(x-1, y)
	left := ha.valueOrMaxInt(x, y+1)
	right := ha.valueOrMaxInt(x, y-1)
	return upper > val && left > val && lower > val && right > val
}

func (ha heightArray) valueOrMaxInt(x, y int) int {
	if ha.isOutOfBounds(x, y) {
		return math.MaxInt
	}
	return ha[x][y].value
}

func (ha heightArray) isBasinEnd(x, y int) bool {
	if ha.isOutOfBounds(x, y) {
		return true
	}
	return ha[x][y].value == 9
}

func (ha heightArray) isOutOfBounds(x, y int) bool {
	return x < 0 || x > len(ha)-1 || y < 0 || y > len(ha[x])-1
}

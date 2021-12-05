package main

import (
	"fmt"

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

	vents := make([][]Point, len(lines))

	for _, l := range lines {
		pAS := l.SubSplitWith(" -> ")

		p1s := pAS[0].SubSplitWith(",")
		p1 := Point{p1s[0].AsInt(), p1s[1].AsInt()}

		p2s := pAS[1].SubSplitWith(",")
		p2 := Point{p2s[0].AsInt(), p2s[1].AsInt()}

		if p1.x == p2.x || p1.y == p2.y {
			vents = append(vents, createVents(p1, p2))
		}
	}
	ventMap := createVentMap(vents)
	return countOverlappingPoints(ventMap)
}

func createVentMap(pts [][]Point) map[Point]int {
	ventMap := make(map[Point]int)
	for _, vts := range pts {
		for _, p := range vts {
			if _, exists := ventMap[p]; !exists {
				ventMap[p] = 0
			}
			ventMap[p]++
		}
	}
	return ventMap
}

func countOverlappingPoints(ventMap map[Point]int) (sum int) {
	for _, v := range ventMap {
		if v > 1 {
			sum++
		}
	}
	return
}

func createVents(start, end Point) []Point {
	res := []Point{}
	fromX, fromY, toX, toY := start.x, start.y, end.x, end.y

	vec := Point{0, 0}
	if fromX > toX {
		vec.x = -1
	}
	if fromX < toX {
		vec.x = 1
	}
	if fromY > toY {
		vec.y = -1
	}
	if fromY < toY {
		vec.y = 1
	}

	for fromY != toY || fromX != toX {
		res = append(res, Point{fromX, fromY})
		fromX += vec.x
		fromY += vec.y
	}
	res = append(res, Point{fromX, fromY})

	return res
}

func partTwo(file util.File) int {
	lines := file.AsLines()

	vents := make([][]Point, len(lines))

	for _, l := range lines {
		pAS := l.SubSplitWith(" -> ")

		p1s := pAS[0].SubSplitWith(",")
		p1 := Point{p1s[0].AsInt(), p1s[1].AsInt()}

		p2s := pAS[1].SubSplitWith(",")
		p2 := Point{p2s[0].AsInt(), p2s[1].AsInt()}

		vents = append(vents, createVents(p1, p2))

	}
	ventMap := createVentMap(vents)
	return countOverlappingPoints(ventMap)
}

package main

import (
	"fmt"

	"github.com/jdrst/adventofgo/util"
)

type VentLocations map[Point]int

type Point struct {
	x, y int
}

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	return GetOverlapping(file, func(p1, p2 Point) bool { return p1.x == p2.x || p1.y == p2.y })
}

func partTwo(file util.File) int {
	return GetOverlapping(file, func(p1, p2 Point) bool { return true })
}

func GetOverlapping(file util.File, constraintFunc func(Point, Point) bool) int {
	lines := file.AsLines()

	vents := make(VentLocations)

	for _, l := range lines {
		pAS := l.SubSplitWith(" -> ")

		p1s := pAS[0].SubSplitWith(",")
		p1 := Point{p1s[0].AsInt(), p1s[1].AsInt()}

		p2s := pAS[1].SubSplitWith(",")
		p2 := Point{p2s[0].AsInt(), p2s[1].AsInt()}

		if constraintFunc(p1, p2) {
			vents.AddVent(p1, p2)
		}
	}
	return vents.OverlappingCount()
}

func (loc VentLocations) OverlappingCount() (sum int) {
	for _, v := range loc {
		if v > 1 {
			sum++
		}
	}
	return
}

func (loc VentLocations) AddPoint(p Point) {
	if _, exists := loc[p]; !exists {
		loc[p] = 0
	}
	loc[p]++
}

func (loc VentLocations) AddVent(start, end Point) {
	delta := Point{0, 0}

	if start.x > end.x {
		delta.x = -1
	}
	if start.x < end.x {
		delta.x = 1
	}
	if start.y > end.y {
		delta.y = -1
	}
	if start.y < end.y {
		delta.y = 1
	}

	for start.y != end.y || start.x != end.x {
		loc.AddPoint(Point{start.x, start.y})
		start.x += delta.x
		start.y += delta.y
	}
	loc.AddPoint(Point{start.x, start.y})
}

package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jdrst/adventofgo/util"
)

type point struct {
	x, y int
}

func main() {
	day13(util.ReadFile("input.txt"))
}

func day13(file util.File) {
	parts := strings.Split(string(file), util.NewLine()+util.NewLine())
	lines := strings.Split(parts[0], util.NewLine())
	folds := strings.Split(parts[1], util.NewLine())
	points := map[point]bool{}

	for _, l := range lines {
		crd := util.Line(l).SubSplitWith(",").AsInts()
		points[point{crd[0], crd[1]}] = true
	}

	var axis string
	var pos, lastX, lastY int
	for i, f := range folds {
		fmt.Sscanf(f, "fold along %1s=%d", &axis, &pos)
		switch axis {
		case "x":
			lastX = pos
		case "y":
			lastY = pos
		}
		foldAlong(axis, pos, points)
		if i == 0 {
			fmt.Printf("First part: %v\n", len(points))
		}
	}

	fmt.Println("Second part:")

	for i := 0; i < lastY; i++ {
		for j := 0; j < lastX; j++ {
			if _, exists := points[point{j, i}]; exists {
				fmt.Printf("%v", "██")
			} else {
				fmt.Printf("%v", "  ")
			}
		}
		fmt.Printf("\n")
	}
}

func foldAlong(axis string, pos int, points map[point]bool) {
	switch axis {
	case "x":
		foldAlongX(pos, points)
	case "y":
		foldAlongY(pos, points)
	default:
		log.Fatal("can't fold")

	}
}

func foldAlongX(pos int, points map[point]bool) {
	for p := range points {
		if p.x > pos {
			delete(points, p)
			points[point{x: pos - (p.x - pos), y: p.y}] = true
		}
	}
}

func foldAlongY(pos int, points map[point]bool) {
	for p := range points {
		if p.y > pos {
			delete(points, p)
			points[point{x: p.x, y: pos - (p.y - pos)}] = true
		}
	}
}

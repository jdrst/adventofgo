package main

import (
	"fmt"

	"github.com/jdrst/adventofgo/util"
)

type node struct {
	p             util.Point
	weight, index int
}

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	cavern := file.AsLines().As2DInts("")

	start := util.Point{X: 0, Y: 0}
	target := util.Point{X: len(cavern) - 1, Y: len(cavern[len(cavern)-1]) - 1}

	return astar(cavern, start, target, func(a, b util.Point) int { return 0 })
}

func partOneManhattanDist(file util.File) int {
	cavern := file.AsLines().As2DInts("")

	start := util.Point{X: 0, Y: 0}
	target := util.Point{X: len(cavern) - 1, Y: len(cavern[len(cavern)-1]) - 1}

	return astar(cavern, start, target, util.ManhattanDistance)
}

func partTwo(file util.File) int {
	cavern := makeCavern(file.AsLines())

	start := util.Point{X: 0, Y: 0}
	target := util.Point{X: len(cavern) - 1, Y: len(cavern[len(cavern)-1]) - 1}

	return astar(cavern, start, target, func(a, b util.Point) int { return 0 })
}

func partTwoManhattanDist(file util.File) int {
	cavern := makeCavern(file.AsLines())

	start := util.Point{X: 0, Y: 0}
	target := util.Point{X: len(cavern) - 1, Y: len(cavern[len(cavern)-1]) - 1}

	return astar(cavern, start, target, util.ManhattanDistance)
}

func makeCavern(lines util.Lines) [][]int {
	cavern := make([][]int, len(lines)*5)

	for i, l := range lines {
		cavern[i+len(lines)*0] = make([]int, len(l)*5)
		cavern[i+len(lines)*1] = make([]int, len(l)*5)
		cavern[i+len(lines)*2] = make([]int, len(l)*5)
		cavern[i+len(lines)*3] = make([]int, len(l)*5)
		cavern[i+len(lines)*4] = make([]int, len(l)*5)
		for j, c := range l.SubSplitWith("").AsInts() {
			cavern[i][j] = c
		}
	}

	height := len(lines)
	width := len(lines[0])
	for multX := 0; multX < 5; multX++ {
		for multY := 0; multY < 5; multY++ {
			if multX == 0 && multY == 0 {
				continue
			}
			for i := height * multY; i < height*(multY+1); i++ {
				for j := width * multX; j < width*(multX+1); j++ {
					newX, newY := i, j
					if j < width {
						newX = i - height
					}
					if j >= width {
						newY = j - width
					}
					val := cavern[newX][newY] + 1
					if val > 9 {
						val = 1
					}
					cavern[i][j] = val
				}
			}
		}
	}

	return cavern
}

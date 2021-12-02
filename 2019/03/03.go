package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strings"

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
	wires := file.AsLines()
	distances, _ := getDistancesAndSteps(wires)
	return distances[0]
}

func partTwo(file util.File) int {
	wires := file.AsLines()
	_, steps := getDistancesAndSteps(wires)
	return steps[0]
}

func getDistancesAndSteps(wires util.Lines) ([]int, []int) {
	wire1 := wireToPointsWithSteps(strings.Split(string(wires[0]), ","))
	wire2 := wireToPointsWithSteps(strings.Split(string(wires[1]), ","))
	distances := make([]int, 0)
	steps := make([]int, 0)
	for p, w1steps := range wire1 {
		if w2steps, exists := wire2[p]; exists {
			distances = append(distances, manhattanDistance(Point{0, 0}, p))
			steps = append(steps, w1steps+w2steps)
		}
	}
	sort.Ints(distances)
	sort.Ints(steps)
	return distances, steps
}

func wireToPointsWithSteps(wire []string) map[Point]int {
	res := make(map[Point]int)
	currPos := Point{0, 0}
	steps := 0
	for _, w := range wire {
		direction := rune(w[0])
		delta := deltaFrom(direction)
		amount := util.ToInt(w[1:])
		for i := 0; i < amount; i++ {
			steps++
			currPos.x += delta.x
			currPos.y += delta.y
			if _, exists := res[currPos]; !exists {
				res[currPos] = steps
			}
		}
	}
	return res
}

func deltaFrom(direction rune) Point {
	switch direction {
	case 'U':
		return Point{0, 1}
	case 'L':
		return Point{-1, 0}
	case 'D':
		return Point{0, -1}
	case 'R':
		return Point{1, 0}
	default:
		log.Fatal("unknown direction")
		return Point{0, 0}
	}
}

func manhattanDistance(p, q Point) int {
	return int(math.Abs(float64(p.x-q.x)) + math.Abs(float64(p.y-q.y)))
}

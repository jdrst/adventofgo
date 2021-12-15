package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/jdrst/adventofgo/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

type node struct {
	x, y, weight int
}

func (n *node) neighbours(maxX, maxY int) []node {
	res := make([]node, 0)
	if n.x > 0 {
		res = append(res, node{n.x - 1, n.y, 0})
	}
	if n.x < maxX {
		res = append(res, node{n.x + 1, n.y, 0})
	}
	if n.y > 0 {
		res = append(res, node{n.x, n.y - 1, 0})
	}
	if n.y < maxY {
		res = append(res, node{n.x, n.y + 1, 0})
	}
	return res
}

func dijkstra(cavern [][]int) int {
	distances := make([][]int, len(cavern))
	for i, l := range cavern {
		distances[i] = make([]int, len(l))
		for j := range l {
			distances[i][j] = math.MaxInt
		}
	}
	seen := make([][]bool, len(cavern))
	for i, l := range cavern {
		seen[i] = make([]bool, len(l))
		for j := range l {
			seen[i][j] = false
		}
	}
	pq := []node{{0, 0, cavern[0][0]}}

	distances[0][0] = cavern[0][0]

	for len(pq) > 0 {
		curr := pq[0]
		if curr.x == len(cavern)-1 && curr.y == len(cavern[len(cavern)-1])-1 {
			break
		}
		pq = pq[1:]
		seen[curr.x][curr.y] = true
		for _, next := range curr.neighbours(len(cavern)-1, len(cavern[curr.x])-1) {
			if seen[next.x][next.y] {
				continue
			}
			if distances[next.x][next.y] > distances[curr.x][curr.y]+cavern[next.x][next.y] {
				distances[next.x][next.y] = distances[curr.x][curr.y] + cavern[next.x][next.y]
				pq = append(pq, node{next.x, next.y, distances[next.x][next.y]})
			}
		}
		sort.Slice(pq, func(i, j int) bool {
			return pq[i].weight < pq[j].weight
		})
	}

	return distances[len(cavern)-1][len(cavern[len(cavern)-1])-1] - cavern[0][0]
}

func partOne(file util.File) int {
	lines := file.AsLines()

	cavern := toCavern(lines)

	return lowestRisk(cavern)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lowestRisk(cavern [][]int) int {

	for i := 1; i < len(cavern[0]); i++ {
		cavern[0][i] += cavern[0][i-1]
	}

	for i := 1; i < len(cavern); i++ {
		cavern[i][0] += cavern[i-1][0]
	}

	for i := 1; i < len(cavern); i++ {
		for j := 1; j < len(cavern[i]); j++ {
			cavern[i][j] += min(cavern[i][j-1], cavern[i-1][j])
		}
	}

	return cavern[len(cavern)-1][len(cavern[len(cavern)-1])-1] - cavern[0][0]
}

func partTwo(file util.File) int {
	lines := file.AsLines()

	cavern := makeCavern(lines)

	return dijkstra(cavern)
}

func toCavern(lines util.Lines) [][]int {
	cavern := make([][]int, len(lines))

	for i, l := range lines {
		cavern[i] = make([]int, len(l))
		for j, c := range l.SubSplitWith("") {
			cavern[i][j] = c.AsInt()
		}
	}
	return cavern
}

func makeCavern(lines util.Lines) [][]int {
	cavern := make([][]int, len(lines)*5)

	for i, l := range lines {
		cavern[i+len(lines)*0] = make([]int, len(l)*5)
		cavern[i+len(lines)*1] = make([]int, len(l)*5)
		cavern[i+len(lines)*2] = make([]int, len(l)*5)
		cavern[i+len(lines)*3] = make([]int, len(l)*5)
		cavern[i+len(lines)*4] = make([]int, len(l)*5)
		for j, c := range l.SubSplitWith("") {
			cavern[i][j] = c.AsInt()
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

package main

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/jdrst/adventofgo/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

type node struct {
	p             util.Point
	weight, index int
}

func dijkstra(cavern [][]int) int {
	distances := make([][]int, len(cavern))
	seen := make([][]bool, len(cavern))
	for i, l := range cavern {
		distances[i] = make([]int, len(l))
		seen[i] = make([]bool, len(l))
		for j := range l {
			distances[i][j] = math.MaxInt
			seen[i][j] = false
		}
	}
	pq := prioQueue{{util.Point{X: 0, Y: 0}, cavern[0][0], 0}}

	distances[0][0] = cavern[0][0]

	heap.Init(&pq)

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*node)
		if curr.p.X == len(cavern)-1 && curr.p.Y == len(cavern[len(cavern)-1])-1 {
			break
		}
		seen[curr.p.X][curr.p.Y] = true
		for _, next := range curr.p.Neighbours(len(cavern)-1, len(cavern[curr.p.X])-1) {
			if seen[next.X][next.Y] {
				continue
			}
			if distances[next.X][next.Y] > distances[curr.p.X][curr.p.Y]+cavern[next.X][next.Y] {
				distances[next.X][next.Y] = distances[curr.p.X][curr.p.Y] + cavern[next.X][next.Y]
				heap.Push(&pq, &node{p: next, weight: distances[next.X][next.Y]})
			}
		}
	}

	return distances[len(cavern)-1][len(cavern[len(cavern)-1])-1] - cavern[0][0]
}

func partOne(file util.File) int {
	lines := file.AsLines()

	cavern := toCavern(lines)

	return dijkstra(cavern)
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

type prioQueue []*node

func (pq prioQueue) Len() int { return len(pq) }

func (pq prioQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

func (pq prioQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *prioQueue) Push(new interface{}) {
	n := len(*pq)
	item := new.(*node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *prioQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

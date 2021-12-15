package main

import (
	"container/heap"
	"math"

	"github.com/jdrst/adventofgo/util"
)

func astar(cavern [][]int, start, target util.Point, heuristicFunc func(a, b util.Point) int) int {
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

	pq := prioQueue{{util.Point{X: 0, Y: 0}, cavern[0][0], heuristicFunc(start, target)}}
	heap.Init(&pq)

	distances[0][0] = cavern[0][0]

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*node)
		if curr.p == target {
			break
		}
		seen[curr.p.X][curr.p.Y] = true
		for _, next := range curr.p.Neighbours(len(cavern)-1, len(cavern[curr.p.X])-1) {
			if seen[next.X][next.Y] {
				continue
			}
			tentative := distances[curr.p.X][curr.p.Y] + cavern[next.X][next.Y]
			if tentative < distances[next.X][next.Y] {
				distances[next.X][next.Y] = tentative
				heap.Push(&pq, &node{p: next, weight: tentative + heuristicFunc(next, target)})
			}
		}
	}

	return distances[target.X][target.Y] - cavern[0][0]
}

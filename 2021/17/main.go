package main

import (
	"fmt"
	"math"

	"github.com/jdrst/adventofgo/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

type trajectory struct {
	xTrajectory, yTrajectory, time int
}

func partOne(file util.File) int {
	lines := file.AsLines()
	var xmin, xmax, ymin, ymax int
	fmt.Sscanf(string(lines[0]), "target area: x=%d..%d, y=%d..%d", &xmin, &xmax, &ymin, &ymax)

	possibleYs := []trajectory{}
next:
	for traj := ymin; traj < util.Abs(ymin); traj++ {
		curr := traj
		if curr <= ymax && curr >= ymin {
			possibleYs = append(possibleYs, trajectory{yTrajectory: traj, time: 0})
		}
		for t := 1; ; t++ {
			curr += traj - t
			if curr < ymin {
				continue next
			}
			if curr <= ymax && curr >= ymin {
				possibleYs = append(possibleYs, trajectory{yTrajectory: traj, time: t})
			}
		}
	}

	maxY := math.MinInt
	for _, t := range possibleYs {
		for traj := 0; traj <= xmax; traj++ {
			curr := traj
			dx := traj
			for i := 0; i < t.time; i++ {
				dx--
				if dx == 0 {
					break
				}
				curr += dx
			}
			if curr <= xmax && curr >= xmin {
				highest := highestPoint(t.yTrajectory)
				if highest > maxY {
					maxY = highest
				}
			}
		}
	}
	return maxY
}

func partTwo(file util.File) int {
	lines := file.AsLines()
	var xmin, xmax, ymin, ymax int
	fmt.Sscanf(string(lines[0]), "target area: x=%d..%d, y=%d..%d", &xmin, &xmax, &ymin, &ymax)

	possibleYs := []trajectory{}
next:
	for traj := ymin; traj < util.Abs(ymin); traj++ {
		curr := traj
		if curr <= ymax && curr >= ymin {
			possibleYs = append(possibleYs, trajectory{yTrajectory: traj, time: 0})
		}
		for t := 1; ; t++ {
			curr += traj - t
			if curr < ymin {
				continue next
			}
			if curr <= ymax && curr >= ymin {
				possibleYs = append(possibleYs, trajectory{yTrajectory: traj, time: t})
			}
		}
	}

	unique := map[util.Point]bool{}
	for _, t := range possibleYs {
		for traj := 0; traj <= xmax; traj++ {
			curr := traj
			dx := traj
			for i := 1; i <= t.time; i++ {
				dx--
				if dx == 0 {
					break
				}
				curr += dx
			}
			if curr <= xmax && curr >= xmin {
				unique[util.Point{X: traj, Y: t.yTrajectory}] = true
			}
		}
	}
	return len(unique)
}

func highestPoint(y int) int {
	return (y * (y + 1)) / 2
}

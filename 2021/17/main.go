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

type foo struct {
	initVelocity, t int
}

type bar struct {
	xVel, yVel, t int
}

func partOne(file util.File) int {
	lines := file.AsLines()
	var xmin, xmax, ymin, ymax int
	fmt.Sscanf(string(lines[0]), "target area: x=%d..%d, y=%d..%d", &xmin, &xmax, &ymin, &ymax)
	xsize := util.Delta(xmax, xmin)
	ysize := util.Delta(ymax, ymin)

	target := make([][]util.Point, xsize+1)
	for i := 0; i <= xsize; i++ {
		target[i] = make([]util.Point, ysize+1)
		for j := 0; j <= ysize; j++ {
			target[i][j] = util.Point{X: xmin + i, Y: ymin + j}
		}
	}

	possibleYs := []foo{}
next:
	for traj := ymin; traj < util.Abs(ymin); traj++ {
		curr := traj
		if curr <= ymax && curr >= ymin {
			possibleYs = append(possibleYs, foo{initVelocity: traj, t: 0})
		}
		for t := 1; ; t++ {
			dy := yy(traj, t)
			curr += dy
			if curr < ymin {
				continue next
			}
			if curr <= ymax && curr >= ymin {
				possibleYs = append(possibleYs, foo{initVelocity: traj, t: t})
			}
		}
	}

	possibleXYs := []bar{}
	for _, t := range possibleYs {
		for traj := 0; traj <= xmax; traj++ {
			curr := traj
			for i := 0; i < t.t; i++ {
				dx := xx(traj, i)
				if dx == 0 {
					break
				}
				curr += dx
			}
			if curr <= xmax && curr >= xmin {
				possibleXYs = append(possibleXYs, bar{xVel: traj, yVel: t.initVelocity, t: t.t})
			}
		}
	}

	maxY := math.MinInt
	for _, xy := range possibleXYs {
		curr := getHighestPointOfTrajectory(xy.yVel)
		if curr > maxY {
			maxY = curr
		}
	}
	return maxY
}

func partTwo(file util.File) int {
	lines := file.AsLines()
	var xmin, xmax, ymin, ymax int
	fmt.Sscanf(string(lines[0]), "target area: x=%d..%d, y=%d..%d", &xmin, &xmax, &ymin, &ymax)
	xsize := util.Delta(xmax, xmin)
	ysize := util.Delta(ymax, ymin)

	target := make([][]util.Point, xsize+1)
	for i := 0; i <= xsize; i++ {
		target[i] = make([]util.Point, ysize+1)
		for j := 0; j <= ysize; j++ {
			target[i][j] = util.Point{X: xmin + i, Y: ymin + j}
		}
	}

	possibleYs := []foo{}
next:
	for traj := ymin; traj < util.Abs(ymin); traj++ {
		curr := traj
		if curr <= ymax && curr >= ymin {
			possibleYs = append(possibleYs, foo{initVelocity: traj, t: 0})
		}
		for t := 1; ; t++ {
			dy := yy(traj, t)
			curr += dy
			if curr < ymin {
				continue next
			}
			if curr <= ymax && curr >= ymin {
				possibleYs = append(possibleYs, foo{initVelocity: traj, t: t})
			}
		}

	}

	possibleXYs := []bar{}
	for _, t := range possibleYs {
		for traj := 0; traj <= xmax; traj++ {
			curr := traj
			for i := 1; i <= t.t; i++ {
				dx := xx(traj, i)
				if dx == 0 {
					break
				}
				curr += dx
			}
			if curr <= xmax && curr >= xmin {
				possibleXYs = append(possibleXYs, bar{xVel: traj, yVel: t.initVelocity, t: t.t})
			}
		}
	}

	unique := map[util.Point]bool{}
	for _, xy := range possibleXYs {
		if _, exists := unique[util.Point{X: xy.xVel, Y: xy.yVel}]; exists {
			continue
		}
		unique[util.Point{X: xy.xVel, Y: xy.yVel}] = true
	}
	return len(unique)
}

func getHighestPointOfTrajectory(y int) int {
	return (y * (y + 1)) / 2
}

func yy(y, t int) int {
	return y - t
}

func xx(x, t int) int {
	for i := 0; i < t; i++ {
		switch {
		case x == 0:
			return x
		case x > 0:
			x--
		case x < 0:
			x++
		}
	}
	return x
}

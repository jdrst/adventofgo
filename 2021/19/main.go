package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/jdrst/adventofgo/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

type vector struct {
	x, y, z int
}

type distance struct {
	from, to vector
	value    int
}

type normalizedVector struct {
	x, y, z float64
}

type scanner struct {
	position vector
	solved   bool
	number   int
	beacons  []vector
	// rotations [][]vector
}

func partOne(file util.File) int {
	input := strings.Split(string(file), util.NewLine()+util.NewLine())

	solvedScanners := solve(input)

	uniqueBeacons := map[vector]bool{}
	for _, ss := range solvedScanners {
		for _, p := range ss.beacons {
			uniqueBeacons[p] = true
		}
	}
	return len(uniqueBeacons)
}

func partTwo(file util.File) int {
	input := strings.Split(string(file), util.NewLine()+util.NewLine())

	solvedScanners := solve(input)

	maxDist := math.MinInt
	{
		for _, s1 := range solvedScanners {
			for _, s2 := range solvedScanners {
				maxDist = max(maxDist, s1.position.taxicab(s2.position).value)
			}
		}
	}

	return maxDist
}

func solve(input []string) []scanner {
	solvedScanners := []scanner{parseScanner(input[0])}
	solvedScanners[0].solved = true
	solvedScanners[0].position = vector{0, 0, 0}
	// solvedScanners[0].rotations = nil

	unsolvedScanners := []scanner{}
	for _, s := range input[1:] {
		unsolvedScanners = append(unsolvedScanners, parseScanner(s))
	}

next:
	for len(unsolvedScanners) > 0 {
		for _, ss := range solvedScanners {
			for i, us := range unsolvedScanners {
				if compareBeacons(&ss, &us) {
					solvedScanners = append(solvedScanners, us)
					if len(unsolvedScanners) > 1 {
						unsolvedScanners = append(unsolvedScanners[:i], unsolvedScanners[i+1:]...)
					} else {
						unsolvedScanners = nil
					}
					goto next
				}
			}
		}
		panic("unsolveable")
	}
	return solvedScanners
}

func parseScanner(s string) scanner {
	lines := strings.Split(s, util.NewLine())
	var num int
	fmt.Sscanf(lines[0], "--- scanner %v ---", &num)
	beacons := make([]vector, len(lines)-1)
	for i, l := range lines[1:] {
		beacons[i] = parseBeacon(l)
	}
	//rotations := getAllRotations(beacons)
	return scanner{number: num, beacons: beacons} //, rotations: rotations}
}

func (s *scanner) absoluteBeaconPositions() []vector {
	res := make([]vector, len(s.beacons))
	for i, b := range s.beacons {
		res[i] = vector{b.x + s.position.x, b.y + s.position.y, b.z + s.position.z}
	}
	return res
}

func compareBeacons(a, b *scanner) bool {
	for _, aBeacon := range a.beacons {
		aDistances := pointDistances(aBeacon, a.beacons)
		for _, bBeacon := range b.beacons {
			bDistances := pointDistances(bBeacon, b.beacons)
			hasEqualDistances, equalADist, equalBDist := compareDistances(aDistances, bDistances)
			if hasEqualDistances {
				a1 := equalADist[0].from.minus(equalADist[0].to) //.normalize()
				bFrom, newBeacons := rotateUntil(a1, equalBDist[0].from, equalBDist[0].to, b.beacons)
				b.beacons = newBeacons
				b.solved = true
				aFrom := equalADist[0].from
				b.position = vector{aFrom.x - bFrom.x, aFrom.y - bFrom.y, aFrom.z - bFrom.z}
				b.beacons = b.absoluteBeaconPositions()
				return true
			}
		}
	}
	return false
}

func pointDistances(p1 vector, points []vector) []distance {
	res := make([]distance, len(points)-1)
	i := 0
	for _, p2 := range points {
		if p1 == p2 {
			continue
		}
		res[i] = p1.taxicab(p2)
		i++
	}
	return res
}

func rotationMatchable(a, b distance) bool {
	v, _ := rotateUntil(a.from.minus(a.to), b.from, b.to, nil)
	nV := vector{}
	if v == nV {
		return false
	}
	return true
}

func compareDistances(a, b []distance) (bool, []distance, []distance) {
	aDistances := []distance{}
	bDistances := []distance{}
	cnt := 0
next:
	for i, dA := range a {
		if cnt > 10 {
			return true, aDistances, bDistances
		}
		for j, dB := range b {
			if dA.value == dB.value && rotationMatchable(dA, dB) {
				aDistances = append(aDistances, dA)
				bDistances = append(bDistances, dB)
				cnt++
				newA := a[:i]
				newB := b[:j]
				if i < len(a) {
					newA = append(newA, a[i+1:]...)
				}
				if j < len(b) {
					newB = append(newB, b[j+1:]...)
				}
				a = newA
				b = newB
				goto next

			}
		}
	}
	return false, nil, nil
}

func parseBeacon(s string) vector {
	coords := strings.Split(s, ",")
	return vector{x: util.ToInt(coords[0]), y: util.ToInt(coords[1]), z: util.ToInt(coords[2])}
}

func rotateUntil(comparison, from, to vector, beacons []vector) (vector, []vector) {
	if from.minus(to) == comparison {
		return from, beacons
	}
	nFrom, nTo, nBeacons := from, to, beacons
	for i := 0; i < 3; i++ {
		nFrom, nTo, nBeacons = rotateY(nFrom, 90), rotateY(nTo, 90), rotatePoints(nBeacons, 90, rotateY)
		if nFrom.minus(nTo) == comparison {
			return nFrom, nBeacons
		}
	}
	nFrom, nTo, nBeacons = rotateX(from, 90), rotateX(to, 90), rotatePoints(beacons, 90, rotateX)
	if nFrom.minus(nTo) == comparison {
		return nFrom, nBeacons
	}
	for i := 0; i < 3; i++ {
		nFrom, nTo, nBeacons = rotateY(nFrom, 90), rotateY(nTo, 90), rotatePoints(nBeacons, 90, rotateY)
		if nFrom.minus(nTo) == comparison {
			return nFrom, nBeacons
		}
	}
	nFrom, nTo, nBeacons = rotateX(from, 270), rotateX(to, 270), rotatePoints(beacons, 270, rotateX)
	if nFrom.minus(nTo) == comparison {
		return nFrom, nBeacons
	}
	for i := 0; i < 3; i++ {
		nFrom, nTo, nBeacons = rotateY(nFrom, 90), rotateY(nTo, 90), rotatePoints(nBeacons, 90, rotateY)
		if nFrom.minus(nTo) == comparison {
			return nFrom, nBeacons
		}
	}
	nFrom, nTo, nBeacons = rotateZ(from, 90), rotateZ(to, 90), rotatePoints(beacons, 90, rotateZ)
	if nFrom.minus(nTo) == comparison {
		return nFrom, nBeacons
	}
	for i := 0; i < 3; i++ {
		nFrom, nTo, nBeacons = rotateY(nFrom, 90), rotateY(nTo, 90), rotatePoints(nBeacons, 90, rotateY)
		if nFrom.minus(nTo) == comparison {
			return nFrom, nBeacons
		}
	}
	nFrom, nTo, nBeacons = rotateZ(from, 180), rotateZ(to, 180), rotatePoints(beacons, 180, rotateZ)
	if nFrom.minus(nTo) == comparison {
		return nFrom, nBeacons
	}
	for i := 0; i < 3; i++ {
		nFrom, nTo, nBeacons = rotateY(nFrom, 90), rotateY(nTo, 90), rotatePoints(nBeacons, 90, rotateY)
		if nFrom.minus(nTo) == comparison {
			return nFrom, nBeacons
		}
	}
	nFrom, nTo, nBeacons = rotateZ(from, 270), rotateZ(to, 270), rotatePoints(beacons, 270, rotateZ)
	if nFrom.minus(nTo) == comparison {
		return nFrom, nBeacons
	}
	for i := 0; i < 3; i++ {
		nFrom, nTo, nBeacons = rotateY(nFrom, 90), rotateY(nTo, 90), rotatePoints(nBeacons, 90, rotateY)
		if nFrom.minus(nTo) == comparison {
			return nFrom, nBeacons
		}
	}
	// panic("no rotations fit")
	return vector{}, nil
}

func rotatePoints(points []vector, degrees int, rotFunc func(vector, int) vector) []vector {
	res := make([]vector, len(points))
	for i, p := range points {
		res[i] = rotFunc(p, degrees)
	}
	return res
}

func rotateX(p vector, degrees int) vector {
	radian := float64(degrees) * (math.Pi / 180)
	cos := math.Round(math.Cos(float64(radian)))
	sin := math.Round(math.Sin(float64(radian)))
	rotM := [3][3]float64{
		{1, 0, 0},
		{0, cos, -sin},
		{0, sin, cos},
	}
	return rotMatMul(p, rotM)
}

func rotateY(p vector, degrees int) vector {
	radian := float64(degrees) * (math.Pi / 180)
	cos := math.Round(math.Cos(float64(radian)))
	sin := math.Round(math.Sin(float64(radian)))
	rotM := [3][3]float64{
		{cos, 0, sin},
		{0, 1, 0},
		{-sin, 0, cos},
	}
	return rotMatMul(p, rotM)
}

func rotateZ(p vector, degrees int) vector {
	radian := float64(degrees) * (math.Pi / 180)
	cos := math.Round(math.Cos(float64(radian)))
	sin := math.Round(math.Sin(float64(radian)))
	rotM := [3][3]float64{
		{cos, -sin, 0},
		{sin, cos, 0},
		{0, 0, 1},
	}
	return rotMatMul(p, rotM)
}

func rotMatMul(p vector, mat [3][3]float64) vector {
	x := float64(p.x)*mat[0][0] + float64(p.y)*mat[0][1] + float64(p.z)*mat[0][2]
	y := float64(p.x)*mat[1][0] + float64(p.y)*mat[1][1] + float64(p.z)*mat[1][2]
	z := float64(p.x)*mat[2][0] + float64(p.y)*mat[2][1] + float64(p.z)*mat[2][2]
	return vector{int(x), int(y), int(z)}
}

func (a vector) minus(b vector) vector {
	return vector{a.x - b.x, a.y - b.y, a.z - b.z}
}

func (a vector) taxicab(b vector) distance {
	return distance{from: a, to: b, value: util.Delta(a.x, b.x) + util.Delta(a.y, b.y) + util.Delta(a.z, b.z)}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

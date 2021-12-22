package main

import (
	"fmt"

	"github.com/jdrst/adventofgo/util"
)

type coord struct {
	x, y, z int
}
type cuboid struct {
	min, max coord
}

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	lines := file.AsLines()

	cubes := []cuboid{}

	var xmin, xmax, ymin, ymax, zmin, zmax int
	var state string
	for _, l := range lines {
		fmt.Sscanf(string(l), "%v x=%d..%d,y=%d..%d,z=%d..%d", &state, &xmin, &xmax, &ymin, &ymax, &zmin, &zmax)
		if xmax < -50 || ymax < -50 || zmax < -50 || xmin > 50 || ymin > 50 || zmin > 50 {
			continue
		}
		xmin = max(xmin, -50)
		xmax = min(xmax, 50)
		ymin = max(ymin, -50)
		ymax = min(ymax, 50)
		zmin = max(zmin, -50)
		zmax = min(zmax, 50)

		min := coord{xmin - 1, ymin - 1, zmin - 1}
		max := coord{xmax, ymax, zmax}
		current := cuboid{min, max}
		newCubes := []cuboid{}
		for _, c := range cubes {
			for _, new := range c.except(current) {
				newCubes = append(newCubes, new)
			}
		}
		if state == "on" {
			newCubes = append(newCubes, current)
		}
		cubes = newCubes
	}

	return sumCubes(cubes)
}

func partTwo(file util.File) int {
	lines := file.AsLines()

	cubes := []cuboid{}

	var xmin, xmax, ymin, ymax, zmin, zmax int
	var state string
	for _, l := range lines {
		fmt.Sscanf(string(l), "%v x=%d..%d,y=%d..%d,z=%d..%d", &state, &xmin, &xmax, &ymin, &ymax, &zmin, &zmax)

		min := coord{xmin - 1, ymin - 1, zmin - 1}
		max := coord{xmax, ymax, zmax}
		current := cuboid{min, max}
		newCubes := []cuboid{}
		for _, c := range cubes {
			for _, new := range c.except(current) {
				newCubes = append(newCubes, new)
			}
		}
		if state == "on" {
			newCubes = append(newCubes, current)
		}
		cubes = newCubes
	}

	return sumCubes(cubes)
}

func (c *cuboid) volume() int {
	return util.Delta(c.max.x, c.min.x) * util.Delta(c.max.y, c.min.y) * util.Delta(c.max.z, c.min.z)
}

func between(num, x, y int) bool {
	if num >= x && num <= y {
		return true
	}
	return false
}

func (a cuboid) intersects(b cuboid) bool {
	return (between(b.min.x, a.min.x, a.max.x) || between(a.min.x, b.min.x, b.max.x)) && (between(b.min.y, a.min.y, a.max.y) || between(a.min.y, b.min.y, b.max.y)) && (between(b.min.z, a.min.z, a.max.z) || between(a.min.z, b.min.z, b.max.z))
}

func (b cuboid) isWithin(a cuboid) bool {
	return between(b.min.x, a.min.x, a.max.x) && between(b.min.y, a.min.y, a.max.y) && between(b.min.z, a.min.z, a.max.z) && between(b.max.x, a.min.x, a.max.x) && between(b.max.y, a.min.y, a.max.y) && between(b.max.z, a.min.z, a.max.z)
}

func (a cuboid) except(b cuboid) []cuboid {
	res := []cuboid{}
	if a.isWithin(b) {
		return res
	}
	if a.intersects(b) {
		minX := min(a.max.x, b.max.x)
		maxX := max(a.min.x, b.min.x)
		//something right
		if minX < a.max.x {
			res = append(res, cuboid{coord{minX, a.min.y, a.min.z}, a.max})
			a.max.x = minX
		}
		//something left
		if maxX > a.min.x {
			res = append(res, cuboid{a.min, coord{maxX, a.max.y, a.max.z}})
			a.min.x = maxX
		}
		minY := min(a.max.y, b.max.y)
		maxY := max(a.min.y, b.min.y)
		//something over
		if minY < a.max.y {
			res = append(res, cuboid{coord{a.min.x, minY, a.min.z}, a.max})
			a.max.y = minY
		}
		//something under
		if maxY > a.min.y {
			res = append(res, cuboid{a.min, coord{a.max.x, maxY, a.max.z}})
			a.min.y = maxY
		}
		minZ := min(a.max.z, b.max.z)
		maxZ := max(a.min.z, b.min.z)
		//something in back
		if minZ < a.max.z {
			res = append(res, cuboid{coord{a.min.x, a.min.y, minZ}, a.max})
			a.max.z = minZ
		}
		//something in front
		if maxZ > a.min.z {
			res = append(res, cuboid{a.min, coord{a.max.x, a.max.y, maxZ}})
			a.min.z = maxZ
		}
		return res
	}
	return []cuboid{a}
}

func sumCubes(cubes []cuboid) int {
	on := 0
	for _, c := range cubes {
		on += c.volume()
	}
	return on
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

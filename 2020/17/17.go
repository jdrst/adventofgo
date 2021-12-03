package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type cubegrid struct {
	maxX, maxY, maxZ, minX, minY, minZ int
	states                             map[cube]bool
}
type hypercubegrid struct {
	maxX, maxY, maxZ, maxW, minX, minY, minZ, minW int
	states                                         map[hypercube]bool
}
type cube struct {
	x, y, z int
}
type hypercube struct {
	x, y, z, w int
}

var newLine = "\r\n"

func main() {
	cubes, hypercubes := prepInput(parseInput())
	cubes.bootCycle()
	cubes.bootCycle()
	cubes.bootCycle()
	cubes.bootCycle()
	cubes.bootCycle()
	cubes.bootCycle()
	activeCubes := 0
	for _, state := range cubes.states {
		if state {
			activeCubes++
		}
	}
	fmt.Println(activeCubes)
	hypercubes.bootCycle()
	hypercubes.bootCycle()
	hypercubes.bootCycle()
	hypercubes.bootCycle()
	hypercubes.bootCycle()
	hypercubes.bootCycle()
	activeHypercubes := 0
	for _, state := range hypercubes.states {
		if state {
			activeHypercubes++
		}
	}
	fmt.Println(activeHypercubes)
}

func (grid *cubegrid) bootCycle() {
	newStates := make(map[cube]bool)

	grid.expand()

	for x := grid.minX; x <= grid.maxX; x++ {
		for y := grid.minY; y <= grid.maxY; y++ {
			for z := grid.minZ; z <= grid.maxZ; z++ {
				currentCube := cube{x, y, z}
				activeAround := grid.activeCubesAround(currentCube)
				if grid.states[currentCube] && activeAround == 2 || activeAround == 3 {
					newStates[currentCube] = true
				}
			}
		}
	}
	grid.states = newStates
}

func (grid cubegrid) activeCubesAround(c cube) int {
	activeAround := 0
	for x := c.x - 1; x <= c.x+1; x++ {
		for y := c.y - 1; y <= c.y+1; y++ {
			for z := c.z - 1; z <= c.z+1; z++ {
				currentCube := cube{x, y, z}
				if state, exists := grid.states[currentCube]; exists && state {
					if currentCube != c {
						activeAround++
					}
				}
			}
		}
	}
	return activeAround
}

func (grid *cubegrid) expand() {
	grid.minX--
	grid.minY--
	grid.minZ--
	grid.maxX++
	grid.maxY++
	grid.maxZ++
}

func (grid *hypercubegrid) bootCycle() {
	newStates := make(map[hypercube]bool)

	grid.expand()

	for x := grid.minX; x <= grid.maxX; x++ {
		for y := grid.minY; y <= grid.maxY; y++ {
			for z := grid.minZ; z <= grid.maxZ; z++ {
				for w := grid.minW; w <= grid.maxW; w++ {
					currentHypercube := hypercube{x, y, z, w}
					activeAround := grid.activeHypercubesAround(currentHypercube)
					if grid.states[currentHypercube] && activeAround == 2 || activeAround == 3 {
						newStates[currentHypercube] = true
					}
				}
			}
		}
	}
	grid.states = newStates
}

func (grid hypercubegrid) activeHypercubesAround(hc hypercube) int {
	activeAround := 0
	for x := hc.x - 1; x <= hc.x+1; x++ {
		for y := hc.y - 1; y <= hc.y+1; y++ {
			for z := hc.z - 1; z <= hc.z+1; z++ {
				for w := hc.w - 1; w <= hc.w+1; w++ {
					currentHypercube := hypercube{x, y, z, w}
					if state, exists := grid.states[currentHypercube]; exists && state {
						if currentHypercube != hc {
							activeAround++
						}
					}
				}
			}
		}
	}
	return activeAround
}

func (grid *hypercubegrid) expand() {
	grid.minX--
	grid.minY--
	grid.minZ--
	grid.minW--
	grid.maxX++
	grid.maxY++
	grid.maxZ++
	grid.maxW++
}

func parseInput() []byte {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return input
}

func prepInput(input []byte) (cubegrid, hypercubegrid) {
	lines := strings.Split(strings.TrimSpace(string(input)), newLine)

	cubes := make(map[cube]bool)
	hypercubes := make(map[hypercube]bool)
	for i, line := range lines {
		states := strings.Split(line, "")
		for j, state := range states {
			if state == "#" {
				cubes[cube{j, i, 0}] = true
				hypercubes[hypercube{j, i, 0, 0}] = true
			} else {
				hypercubes[hypercube{j, i, 0, 0}] = false
				cubes[cube{j, i, 0}] = false
			}
		}
	}
	cubegrid := cubegrid{maxX: len(lines[0]) - 1, maxY: len(lines) - 1, maxZ: 0, minX: 0, minY: 0, minZ: 0, states: cubes}
	hypercubegrid := hypercubegrid{maxX: len(lines[0]) - 1, maxY: len(lines) - 1, maxZ: 0, maxW: 0, minX: 0, minY: 0, minZ: 0, minW: 0, states: hypercubes}
	return cubegrid, hypercubegrid
}

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

type amphipod struct {
	tpe                        rune
	energyPerStep, destination int
	x, y                       int
	hasMoved                   bool
}

type burrow struct {
	depth      int
	usedEnergy int
	pods       map[util.Point]amphipod
}

type cache struct {
	currentBest *int
	results     map[string]int
}

func partOne(file util.File) int {
	lines := file.AsLines()

	max := math.MaxInt
	c := cache{currentBest: &max, results: map[string]int{}}
	res := moveUntilDone(makeBurrow(lines), &c)
	return res
}

func makeBurrow(lines util.Lines) burrow {
	pods := lines[2 : len(lines)-1]
	return burrow{usedEnergy: 0, pods: parseAmphipods(pods), depth: len(pods)}
}

func moveUntilDone(b burrow, c *cache) int {
	// fmt.Println(b)
	for _, p := range b.pods {
		if p.x != p.destination {
			goto notDone
		}
	}
	return b.usedEnergy
notDone:

	if v, exists := c.results[fmt.Sprint(b)]; exists {
		return v
	}
	if b.usedEnergy > *c.currentBest {
		return math.MaxInt
	}
	res := math.MaxInt
	for _, p := range b.pods {
		for _, next := range p.move(b) {
			energyUsed := moveUntilDone(next, c)
			res = min(res, energyUsed)
		}
	}

	if res < *c.currentBest {
		c.currentBest = &res
	}
	c.results[fmt.Sprint(b)] = res
	return res
}

func (a amphipod) move(b burrow) (possibilities []burrow) {
	energy := b.usedEnergy
	//remove self
	currentMap := copyMap(b.pods)
	delete(currentMap, util.Point{X: a.x, Y: a.y})
	//move out of side room
	if a.y > 0 {
		if a.hasMoved {
			return
		}
		//won't move if in correct destination
		if a.y < 4 && a.x == a.destination {
			for y := 4; y >= a.y; y-- {
				if p, exists := currentMap[util.Point{X: a.x, Y: y}]; exists {
					if p.tpe != a.tpe {
						goto notDone
					}
				}
			}
			return
		}
	notDone:
		//can't move at all if entry is blocked
		for y := a.y - 1; y >= 0; y-- {
			if _, exists := currentMap[util.Point{X: a.x, Y: y}]; exists {
				return
			}
		}

		_, blockedLeft := currentMap[util.Point{X: a.x - 1, Y: 0}]
		_, blockedRight := currentMap[util.Point{X: a.x + 1, Y: 0}]
		if blockedLeft && blockedRight {
			return
		}
		//move upwards
		newY := a.y
		for newY > 0 {
			newY--
			energy += a.energyPerStep
		}

		//left possibilities
		var newEnergy int
		if !blockedLeft {
			newEnergy = energy
			for x := a.x - 1; x >= 0; x-- {
				newEnergy += a.energyPerStep
				if x == 2 || x == 4 || x == 6 || x == 8 {
					continue
				}
				if _, exists := currentMap[util.Point{X: x, Y: 0}]; exists {
					break
				}
				tempPods := copyMap(currentMap)
				tempPods[util.Point{X: x, Y: newY}] = amphipod{tpe: a.tpe, energyPerStep: a.energyPerStep, x: x, y: newY, destination: a.destination, hasMoved: true}
				possibilities = append(possibilities, burrow{usedEnergy: newEnergy, pods: tempPods, depth: b.depth})

			}
		}
		//right possibilities
		if !blockedRight {
			newEnergy = energy
			for x := a.x + 1; x < 11; x++ {
				newEnergy += a.energyPerStep
				if x == 2 || x == 4 || x == 6 || x == 8 {
					continue
				}
				if _, exists := currentMap[util.Point{X: x, Y: 0}]; exists {
					break
				}
				tempPods := copyMap(currentMap)
				tempPods[util.Point{X: x, Y: newY}] = amphipod{tpe: a.tpe, energyPerStep: a.energyPerStep, x: x, y: newY, destination: a.destination, hasMoved: true}
				possibilities = append(possibilities, burrow{usedEnergy: newEnergy, pods: tempPods, depth: b.depth})
			}
		}
	}
	//move into side room
	if a.y == 0 {
		//can't move if side room occupated
		if _, exists := currentMap[util.Point{X: a.destination, Y: 1}]; exists {
			return
		}
		//don't move in if occupied with wrong type
		for y := b.depth; y > 0; y-- {
			if p, exists := currentMap[util.Point{X: a.destination, Y: y}]; exists {
				if p.tpe != a.tpe {
					return
				}
			}
		}
		//can't move if pod is in the way
		if a.x > a.destination {
			for x := a.x - 1; x >= a.destination; x-- {
				if _, exists := currentMap[util.Point{X: x, Y: 0}]; exists {
					return
				}
			}
		}
		if a.x < a.destination {
			for x := a.x + 1; x <= a.destination; x++ {
				if _, exists := currentMap[util.Point{X: x, Y: 0}]; exists {
					return
				}
			}
		}
		//move into target
		newY := b.depth
		for newY > 0 {
			if _, exists := currentMap[util.Point{X: a.destination, Y: newY}]; !exists {
				break
			}
			newY--
		}
		newPods := copyMap(currentMap)
		newEnergy := energy + util.Delta(a.x, a.destination)*a.energyPerStep + newY*a.energyPerStep
		newPods[util.Point{X: a.destination, Y: newY}] = amphipod{tpe: a.tpe, energyPerStep: a.energyPerStep, x: a.destination, y: newY, destination: a.destination, hasMoved: true}
		possibilities = append(possibilities, burrow{usedEnergy: newEnergy, pods: newPods, depth: b.depth})
	}
	return possibilities
}

func copyMap(m map[util.Point]amphipod) map[util.Point]amphipod {
	res := map[util.Point]amphipod{}
	for k, v := range m {
		res[k] = v
	}
	return res
}

func parseAmphipods(lines util.Lines) map[util.Point]amphipod {
	burrow := map[util.Point]amphipod{}
	for i, l := range lines {
		for j, c := range string(l) {
			var ePs, dest int
			switch c {
			case 'A':
				ePs = 1
				dest = 2
			case 'B':
				ePs = 10
				dest = 4
			case 'C':
				ePs = 100
				dest = 6
			case 'D':
				ePs = 1000
				dest = 8
			default:
				continue
			}
			burrow[util.Point{X: j - 1, Y: i + 1}] = amphipod{tpe: c, energyPerStep: ePs, destination: dest, x: j - 1, y: i + 1}
		}
	}
	return burrow
}

func (b burrow) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("used energy: %v\n", b.usedEnergy))
	sb.WriteString("#############\n")
	sb.WriteRune('#')
	for x, y := 0, 0; x < 11; x++ {
		if p, exists := b.pods[util.Point{X: x, Y: y}]; exists {
			sb.WriteRune(p.tpe)
		} else {
			sb.WriteRune('.')
		}
	}
	sb.WriteString("#\n")
	for y := 1; y <= b.depth; y++ {
		if y == 1 {
			sb.WriteString("##")
		} else {
			sb.WriteString("  ")
		}
		for x := 1; x < 10; x++ {
			if x%2 == 1 {
				sb.WriteRune('#')
				continue
			}
			if p, exists := b.pods[util.Point{X: x, Y: y}]; exists {
				sb.WriteRune(p.tpe)
			} else {
				sb.WriteRune('.')
			}
		}
		if y == 1 {
			sb.WriteString("##")
		}
		sb.WriteString("\n")
	}
	sb.WriteString("  #########")
	return sb.String()
}

func partTwo(file util.File) int {
	lines := file.AsLines()

	lines = append(lines[:5], lines[3:]...)
	lines[3] = "  #D#C#B#A#"
	lines[4] = "  #D#B#A#C#"
	max := math.MaxInt
	c := cache{currentBest: &max, results: map[string]int{}}
	res := moveUntilDone(makeBurrow(lines), &c)
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

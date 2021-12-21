package main

import (
	"fmt"

	"github.com/jdrst/adventofgo/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

type player struct {
	space, points int
}
type deterministicDie struct {
	side  int
	rolls int
}

type diracDieGame struct {
	cache map[[4]int][2]int
}

func partOne(file util.File) int {
	lines := file.AsLines()
	var p1space, p2space int
	fmt.Sscanf(string(lines[0]), "Player 1 starting position: %v", &p1space)
	fmt.Sscanf(string(lines[1]), "Player 2 starting position: %v", &p2space)
	playerOne := player{space: p1space, points: 0}
	playerTwo := player{space: p2space, points: 0}
	p1Turn := true
	die := deterministicDie{0, 0}
	for playerOne.points < 1000 && playerTwo.points < 1000 {
		roll := die.roll() + die.roll() + die.roll()
		if p1Turn {
			playerOne = playerOne.move(roll)
		} else {
			playerTwo = playerTwo.move(roll)
		}
		p1Turn = !p1Turn
	}

	if playerOne.points > playerTwo.points {
		return playerTwo.points * die.rolls
	}
	return playerOne.points * die.rolls
}

func partTwo(file util.File) int {
	lines := file.AsLines()
	var p1space, p2space int
	fmt.Sscanf(string(lines[0]), "Player 1 starting position: %v", &p1space)
	fmt.Sscanf(string(lines[1]), "Player 2 starting position: %v", &p2space)

	game := diracDieGame{cache: map[[4]int][2]int{}}
	wins := game.play(p1space-1, p2space-1, 0, 0)

	if wins[0] > wins[1] {
		return wins[0]
	}
	return wins[1]
}

func (p player) move(steps int) player {
	p.space = (p.space + steps) % 10
	if p.space == 0 {
		p.space = 10
	}
	p.points += p.space
	return p
}

func (d *deterministicDie) roll() int {
	d.rolls++
	d.side++
	if d.side > 100 {
		d.side = 1
	}
	return d.side
}

func (d *diracDieGame) play(currentPos, otherPos, currentPoints, otherPoints int) [2]int {
	if currentPoints > 20 {
		return [2]int{1, 0}
	}
	if otherPoints > 20 {
		return [2]int{0, 1}
	}

	if res, exists := d.cache[[4]int{currentPos, otherPos, currentPoints, otherPoints}]; exists {
		return res
	}

	res := [2]int{0, 0}
	for roll, additional := range quantumDieOutcomes {
		newPos := (currentPos + roll) % 10
		newPoints := currentPoints + newPos + 1
		wins := d.play(otherPos, newPos, otherPoints, newPoints)
		res[0] += wins[1] * additional
		res[1] += wins[0] * additional
	}

	d.cache[[4]int{currentPos, otherPos, currentPoints, otherPoints}] = res
	return res
}

var quantumDieOutcomes map[int]int = map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}

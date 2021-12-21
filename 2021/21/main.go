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
	playerOne := player{p1space, 0}
	playerTwo := player{p2space, 0}

	p1wins, p2wins := quantumDieGame(playerOne, playerTwo, true, 1)

	if p1wins > p2wins {
		return p1wins
	}
	return p2wins
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

func quantumDieGame(current, other player, p1Turn bool, universes int) (p1win int, p2win int) {
	if other.points > 20 {
		if p1Turn {
			return 0, universes
		}
		return universes, 0
	}

	for roll, additional := range quantumDieOutcomes {
		p1wins, p2wins := quantumDieGame(other, current.move(roll), !p1Turn, universes*additional)
		p1win += p1wins
		p2win += p2wins
	}

	return p1win, p2win
}

var quantumDieOutcomes map[int]int = map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}

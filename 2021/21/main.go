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

type quantumPlayer struct {
	space, points, universes int
}

type deterministicDie struct {
	side  int
	rolls int
}

func partOne(file util.File) int {
	lines := file.AsLines()
	var p1, p1space, p2, p2space int
	fmt.Sscanf(string(lines[0]), "Player %v starting position: %v", &p1, &p1space)
	fmt.Sscanf(string(lines[1]), "Player %v starting position: %v", &p2, &p2space)
	playerOne := player{space: p1space, points: 0}
	playerTwo := player{space: p2space, points: 0}
	currentPlayer := &playerOne
	die := deterministicDie{0, 0}
	for playerOne.points < 1000 && playerTwo.points < 1000 {
		roll := die.roll() + die.roll() + die.roll()
		currentPlayer.move(roll, 0)
		if *currentPlayer == playerOne {
			currentPlayer = &playerTwo
		} else {
			currentPlayer = &playerOne
		}
	}

	if playerOne.points > playerTwo.points {
		return playerTwo.points * die.rolls
	}
	return playerOne.points * die.rolls
}

func partTwo(file util.File) int {
	lines := file.AsLines()
	var p1, p1space, p2, p2space int
	fmt.Sscanf(string(lines[0]), "Player %v starting position: %v", &p1, &p1space)
	fmt.Sscanf(string(lines[1]), "Player %v starting position: %v", &p2, &p2space)
	playerOne := quantumPlayer{p1space, 0, 1}
	playerTwo := quantumPlayer{p2space, 0, 1}

	p1wins, p2wins := quantumDieGame(playerOne, playerTwo, true)

	if p1wins > p2wins {
		return p1wins
	}
	return p2wins
}

func (p *player) move(steps int, universes int) player {
	p.space = (p.space + steps) % 10
	if p.space == 0 {
		p.space = 10
	}
	p.points += p.space
	return *p
}

func (p quantumPlayer) move(steps int, universes int) quantumPlayer {
	p.space = (p.space + steps) % 10
	if p.space == 0 {
		p.space = 10
	}
	p.points += p.space
	p.universes *= universes
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

func quantumDieGame(current, other quantumPlayer, p1Turn bool) (p1win int, p2win int) {
	if other.points > 20 {
		if p1Turn {
			return 0, other.universes * current.universes
		}
		return other.universes * current.universes, 0
	}

	for roll, universes := range quantumDieOutcomes {
		p1wins, p2wins := quantumDieGame(other, current.move(roll, universes), !p1Turn)
		p1win += p1wins
		p2win += p2wins
	}

	return p1win, p2win
}

var quantumDieOutcomes map[int]int = map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}

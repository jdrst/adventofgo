package main

import (
	"fmt"
	"strings"

	"github.com/jdrst/adventofgo/util"
)

type BingoNumber struct {
	num    int
	marked bool
}

type Board [5][5]BingoNumber

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	lines := strings.Split(strings.TrimSpace(string(file)), util.NewLine()+util.NewLine())
	nums := util.Line(lines[0]).SubSplitWith(",").AsInts()
	boards := createBoardsFrom(lines[1:])
	for _, n := range nums {
		for j := range boards {
			boards[j].MarkNum(n)
			if boards[j].HasWon() {
				return boards[j].CalculateScore(n)
			}
		}
	}
	return 0
}

func partTwo(file util.File) int {
	lines := strings.Split(strings.TrimSpace(string(file)), util.NewLine()+util.NewLine())
	nums := util.Line(lines[0]).SubSplitWith(",").AsInts()
	boards := createBoardsFrom(lines[1:])

	var lastWinningScore int

	for _, n := range nums {
		for j := 0; j < len(boards); j++ {
			boards[j].MarkNum(n)
			if boards[j].HasWon() {
				lastWinningScore = boards[j].CalculateScore(n)
				boards = append(boards[:j], boards[j+1:]...)
				j--
			}
		}
	}
	return lastWinningScore
}

func (b *Board) CalculateScore(winNr int) int {
	sum := 0
	for _, r := range b {
		for _, n := range r {
			if !n.marked {
				sum += n.num
			}
		}
	}
	return sum * winNr
}

func (b *Board) MarkNum(num int) {
	for i, r := range b {
		for j, n := range r {
			if n.num == num {
				b[i][j].marked = true
			}
		}
	}
}

func (b Board) HasWon() bool {

nextRow:
	for i, r := range b {
		for _, n := range r {
			if !n.marked {
				goto cols
			}
		}
		return true
	cols:
		for j := 0; j < len(b); j++ {
			if !b[j][i].marked {
				continue nextRow
			}
		}
		return true
	}
	return false
}

func createBoardsFrom(a []string) []Board {
	boards := make([]Board, len(a))

	for i, s := range a {
		var b Board
		for i, r := range strings.Split(s, util.NewLine()) {
			for j, n := range strings.Fields(r) {
				b[i][j] = BingoNumber{util.ToInt(n), false}
			}
		}
		boards[i] = b
	}

	return boards
}

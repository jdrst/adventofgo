package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/jdrst/adventofgo/util"
)

func newLine() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}

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
	lines := strings.Split(strings.TrimSpace(string(file)), newLine()+newLine())
	nums := util.Line(lines[0]).SubSplitWith(",").AsInts()
	stringBoards := lines[1:]
	boards := make([]Board, len(stringBoards))
	for i, b := range stringBoards {
		boards[i] = createBoardFrom(b)
	}
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

func createBoardFrom(s string) Board {
	var b Board
	for i, r := range strings.Split(s, newLine()) {
		for j, n := range strings.Fields(r) {
			b[i][j] = BingoNumber{util.ToInt(n), false}
		}
	}
	return b
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
	for _, r := range b {
		cnt := 0
		for _, n := range r {
			if n.marked {
				cnt++
			}
		}
		if cnt == 5 {
			return true
		}
	}
	for i := 0; i < 5; i++ {
		cnt := 0
		for j := 0; j < 5; j++ {
			if b[j][i].marked {
				cnt++
			}
		}
		if cnt == 5 {
			return true
		}
	}
	return false
}

func partTwo(file util.File) int {
	lines := strings.Split(strings.TrimSpace(string(file)), newLine()+newLine())
	nums := util.Line(lines[0]).SubSplitWith(",").AsInts()
	stringBoards := lines[1:]
	boards := make([]Board, len(stringBoards))

	var lastWinningScore int

	for i, b := range stringBoards {
		boards[i] = createBoardFrom(b)
	}
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

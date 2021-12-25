package main

import (
	"fmt"

	"github.com/jdrst/adventofgo/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

type instruction struct {
	a, b, c int
}

func partOne(file util.File) int {
	lines := file.AsLines()

	return solve(parseInstructions(lines), []int{9, 8, 7, 6, 5, 4, 3, 2, 1})
}

func partTwo(file util.File) int {
	lines := file.AsLines()

	return solve(parseInstructions(lines), []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func solve(instructions [14]instruction, order []int) int {
	sol := [14]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	stack := []int{}
	funcs := [14](func(int) int){}
	setFuncIdxs := []int{}
	for i, in := range instructions {
		switch in.a {
		case 1:
			c := in.c
			fn := func(n int) int { return n + c }
			funcs[i] = fn
			stack = append(stack, i)
		case 26:
			fnIdx := stack[len(stack)-1]
			currentIdx := i
			b := in.b
			setFunc := func(n int) int {
				sol[fnIdx] = n
				sol[currentIdx] = funcs[fnIdx](n) + b
				return 0
			}
			setFuncIdxs = append(setFuncIdxs, currentIdx)
			stack = stack[:len(stack)-1]
			funcs[i] = setFunc
		}
	}

	for _, idx := range setFuncIdxs {
		for _, n := range order {
			funcs[idx](n)
			if isValid(sol) {
				break
			}
		}
	}

	return toInt(sol)
}

func isValid(sol [14]int) bool {
	for _, v := range sol {
		if v < 1 || v > 9 {
			return false
		}
	}
	return true
}

func toInt(sol [14]int) int {
	res := 0
	for _, v := range sol {
		res *= 10
		res += v
	}
	return res
}

func parseInstructions(lines util.Lines) [14]instruction {
	res := [14]instruction{}
	aLnNum, bLnNum, cLnNum := 4, 5, 15
	for i := 0; i < 14; i++ {
		aLn := lines[aLnNum].SubSplitWith(" ")
		bLn := lines[bLnNum].SubSplitWith(" ")
		cLn := lines[cLnNum].SubSplitWith(" ")
		res[i] = instruction{a: aLn[2].AsInt(), b: bLn[2].AsInt(), c: cLn[2].AsInt()}
		aLnNum += 18
		bLnNum += 18
		cLnNum += 18
	}
	return res
}

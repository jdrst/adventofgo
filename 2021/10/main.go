package main

import (
	"fmt"
	"sort"

	"github.com/jdrst/adventofgo/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

type Stack []rune

func partOne(file util.File) int {
	lines := file.AsLines()
	stack := make(Stack, 0)
	score := 0
	syntaxerrorpoints := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}

next:
	for _, l := range lines {
		for _, c := range l {
			switch c {
			case '(', '[', '{', '<':
				stack.push(c)
			case ')', ']', '}', '>':
				last := stack.pop()
				if !isCorrectClosure(last, c) {
					score += syntaxerrorpoints[c]
					stack.clear()
					continue next
				}
			}
		}
	}
	return score
}

func partTwo(file util.File) int {
	lines := file.AsLines()
	stack := make(Stack, 0)
	score := make([]int, 0)
	autocompletepoints := map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}

next:
	for _, l := range lines {
		for _, c := range l {
			switch c {
			case '(', '[', '{', '<':
				stack.push(c)
			case ')', ']', '}', '>':
				last := stack.pop()
				if !isCorrectClosure(last, c) {
					stack.clear()
					continue next
				}
			}
		}
		currentscore := 0
		for len(stack) > 0 {
			currentscore *= 5
			currentscore += autocompletepoints[stack.pop()]
		}
		score = append(score, currentscore)
	}
	sort.Slice(score, func(i, j int) bool { return score[i] < score[j] })
	return score[len(score)/2]
}

func (s *Stack) push(r rune) {
	*s = append(*s, r)
}

func (s *Stack) pop() rune {
	if len(*s) == 0 {
		return ' '
	}
	last := []rune(*s)[len(*s)-1]
	*s = []rune(*s)[:len(*s)-1]
	return last
}

func (s *Stack) clear() {
	*s = make(Stack, 0)
}

func isCorrectClosure(open, close rune) bool {
	switch open {
	case '(':
		return close == ')'
	case '[':
		return close == ']'
	case '{':
		return close == '}'
	case '<':
		return close == '>'
	default:
		return false
	}
}

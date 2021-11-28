package main

import (
	"fmt"
	"strings"

	"git.threeman.info/jd/adventofcode/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	lines := file.AsLines()
	bounds := strings.Split(string(lines[0]), "-")
	amount := 0
	for i := util.ToInt(bounds[0]); i <= util.ToInt(bounds[1]); i++ {
		d := toDigits(i)
		if hasDecreasingDigits(d) {
			continue
		}
		if hasSameDigitAdjacent(d) {
			amount++
		}
	}
	return amount
}

func hasDecreasingDigits(digits []rune) bool {
	current := util.ToInt(string(digits[0]))
	for _, d := range digits[1:] {
		next := util.ToInt(string(d))
		if current > next {
			return true
		}
		current = next
	}
	return false
}

func toDigits(num int) []rune {
	res := make([]rune, 0)
	for num > 0 {
		d := num % 10
		res = append([]rune{rune('0' + d)}, res...)
		num = num / 10
	}
	return res
}

func hasSameDigitAdjacent(digits []rune) bool {
	current := digits[0]
	for _, d := range digits[1:] {
		next := d
		if current == next {
			return true
		}
		current = next
	}
	return false
}

func hasSameDigitAdjacentWithoutLargerGroup(digits []rune) bool {
	current := digits[0]
	for _, d := range digits[1:] {
		next := d
		if current == next {
			if count(digits, current) == 2 {
				return true
			}
		}
		current = next
	}
	return false
}

func partTwo(file util.File) int {
	lines := file.AsLines()
	bounds := strings.Split(string(lines[0]), "-")
	amount := 0
	for i := util.ToInt(bounds[0]); i <= util.ToInt(bounds[1]); i++ {
		d := toDigits(i)
		if hasDecreasingDigits(d) {
			continue
		}
		if hasSameDigitAdjacentWithoutLargerGroup(d) {
			amount++
		}
	}
	return amount
}

func count(digits []rune, c rune) int {
	res := 0
	for _, d := range digits {
		if d == c {
			res++
		}
	}
	return res
}

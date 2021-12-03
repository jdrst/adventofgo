package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func parseInput() []byte {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return input
}

var newLine = "\r\n"

func prepInput(input []byte) []string {
	lines := strings.Split(strings.TrimSpace(string(input)), newLine)

	return lines
}

func main() {
	lines := prepInput(parseInput())

	part1, part2 := 0, 0
	re := regexp.MustCompile(`\([^\(\)]+\)`)
	for _, line := range lines {
		part1 += resolveExprFunc(line, samePrecedence, re)
		part2 += resolveExprFunc(line, inversePrecedence, re)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func resolveExprFunc(line string, fn func(string) int, re *regexp.Regexp) int {
	for re.MatchString(line) {
		line = re.ReplaceAllStringFunc(line, func(line string) string { return strconv.Itoa(fn(line)) })
	}
	return fn(line)
}

func samePrecedence(expr string) int {
	parts := strings.Fields(strings.Trim(expr, "()"))
	res, _ := strconv.Atoi(parts[0])

	for i := 1; i < len(parts); i += 2 {
		val, _ := strconv.Atoi(parts[i+1])
		switch parts[i] {
		case "+":
			res += val
		case "*":
			res *= val
		}
	}
	return res
}

func inversePrecedence(expr string) int {
	return resolveExprFunc(expr, samePrecedence, regexp.MustCompile(`\d+ \+ \d+`))
}

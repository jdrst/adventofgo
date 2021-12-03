package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type operation struct {
	order    string
	argument int
}

func parseInput() []operation {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\r\n")
	operations := make([]operation, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)
		arg, _ := strconv.Atoi(fields[1])
		operations[i] = operation{fields[0], arg}
	}
	return operations
}

func main() {
	operations := parseInput()

	part1, _ := checkForLoop(operations)
	fmt.Println(part1)
	fmt.Println(part2(operations))
}

func part2(operations []operation) int {
	for i, op := range operations {
		if op.order == "nop" {
			operations[i] = operation{
				order:    "jmp",
				argument: op.argument,
			}
		} else if op.order == "jmp" {
			operations[i] = operation{
				order:    "nop",
				argument: op.argument,
			}
		}
		acc, done := checkForLoop(operations)
		if done {
			return acc
		}
		operations[i] = op
	}
	return -1
}

func checkForLoop(operations []operation) (int, bool) {
	opsExec := make(map[int]struct{})
	acc := 0
	for i := 0; i < len(operations); i++ {
		if _, exists := opsExec[i]; exists {
			return acc, false
		}
		opsExec[i] = struct{}{}
		operation := operations[i]
		switch operation.order {
		case "acc":
			acc += operation.argument
		case "jmp":
			i += operation.argument - 1
		default:
		}
	}
	return acc, true
}

package main

import (
	"fmt"
	"io/ioutil"
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

func prepInput(input []byte) map[int]int {
	lines := strings.Split(strings.TrimSpace(string(input)), newLine)
	numbers := make(map[int]int)
	for i, numStr := range strings.Split(lines[0], ",") {
		num, _ := strconv.Atoi(numStr)
		numbers[num] = i + 1
	}

	return numbers
}

func main() {
	numbers := prepInput(parseInput())
	fmt.Println(numberN(numbers, 16, 2020))

	numbers = prepInput(parseInput())
	fmt.Println(numberN(numbers, 16, 30000000))
}

func numberN(numbers map[int]int, numSpoken, rounds int) int {
	for i := len(numbers); i < rounds; i++ {
		if spokenLast, exists := numbers[numSpoken]; exists {
			numbers[numSpoken], numSpoken = i, i-spokenLast
		} else {
			numbers[numSpoken], numSpoken = i, 0
		}
	}
	return numSpoken
}

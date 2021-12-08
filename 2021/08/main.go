package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jdrst/adventofgo/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	lines := file.AsLines()

	sublines := make([][]util.Line, len(lines))

	sum := 0

	for i, l := range lines {
		sublines[i] = l.SubSplitWith(" | ")
		digits := strings.Fields(string(sublines[i][1]))
		for _, d := range digits {
			if len(d) == 2 || len(d) == 3 || len(d) == 4 || len(d) == 7 {
				sum++
			}
		}
	}
	return sum
}

func partTwo(file util.File) int {
	lines := file.AsLines()
	sum := 0

	for _, l := range lines {
		split := l.SubSplitWith(" | ")
		sum += deduceValue(string(split[0]), string(split[1]))
	}
	return sum
}

func deduceValue(pattern, output string) int {
	digits := make(map[int]string)
	known := [10]int{}
	patterns := strings.Fields(pattern)
	for _, p := range patterns {
		if len(p) == 2 {
			known[1] = sum(p)
		}
		if len(p) == 4 {
			known[4] = sum(p)
		}
		if len(p) == 3 {
			known[7] = sum(p)
		}
		if len(p) == 7 {
			known[8] = sum(p)
		}
	}
	digits[known[7]-known[1]] = "a"

	counts := map[rune]int{'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0}
	for _, r := range "abcdefg" {
		for _, s := range patterns {
			if strings.ContainsRune(s, r) {
				counts[r]++
			}
		}
	}

	for k, v := range counts {
		if v == 4 {
			digits[int(k)] = "e"
		}
	}

	digits[known[8]-known[4]-sumOfKeys(digits)] = "g"

	for k, v := range counts {
		if v == 9 {
			digits[int(k)] = "f"
			digits[known[1]-int(k)] = "c"
		}
		if v == 6 {
			digits[int(k)] = "b"
		}
	}

	digits[known[8]-sumOfKeys(digits)] = "d"

	outputs := strings.Fields(output)
	mappedOutputs := make([]string, len(outputs))
	for i, s := range outputs {
		mappedOutputs[i] = getDigit(s, digits)
	}

	return getInt(mappedOutputs)
}

func getDigit(s string, mapping map[int]string) string {
	result := ""
	for _, c := range s {
		result += mapping[int(c)]
	}
	return result
}

func sumOfKeys(known map[int]string) int {
	sum := 0
	for k := range known {
		sum += k
	}
	return sum
}
func sum(s string) int {
	sum := 0
	for _, r := range s {
		sum += int(r)
	}
	return sum
}

func getInt(input []string) int {
	var displayDigits map[string]string = map[string]string{"abcefg": "0", "cf": "1", "acdeg": "2", "acdfg": "3", "bcdf": "4", "abdfg": "5", "abdefg": "6", "acf": "7", "abcdefg": "8", "abcdfg": "9"}
	stringResult := ""
next:
	for _, s := range input {
		for k, v := range displayDigits {
			if IsInAnyOrder(s, k) {
				stringResult += v
				continue next
			}
		}
	}
	res, err := strconv.Atoi(stringResult)
	if err != nil {
		log.Fatal("getInt can't convert")
	}
	return res
}

func IsInAnyOrder(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	result := true
next:
	for _, cA := range a {
		for _, cB := range b {
			if cA == cB {
				continue next
			}
		}
		result = false
	}
	return result
}

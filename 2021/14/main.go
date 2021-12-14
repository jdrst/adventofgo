package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/jdrst/adventofgo/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	return polymerizate(file, 10)
}

func polymerizate(file util.File, steps int) int {
	lines := strings.Split(string(file), util.NewLine()+util.NewLine())
	template := lines[0]

	rules := map[string]string{}

	for _, r := range util.File(lines[1]).AsLines() {
		var pair, res string
		fmt.Sscanf(string(r), "%2v -> %v", &pair, &res)
		rules[pair] = res
	}

	pairs := map[string]int{}

	for i := 0; i < len(template)-1; i++ {
		pairs[string(template[i])+string(template[i+1])]++
	}

	for i := 0; i < steps; i++ {
		newPairs := map[string]int{}
		for k, v := range pairs {
			newPairs[string(k[0])+rules[k]] += v
			newPairs[rules[k]+string(k[1])] += v
		}
		pairs = newPairs
	}

	counts := map[rune]int{}
	for k, v := range pairs {
		counts[rune(k[0])] += v
		counts[rune(k[1])] += v
	}

	max, min := math.MinInt, math.MaxInt

	for _, v := range counts {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return (max-min)/2 + 1
}

func partTwo(file util.File) int {
	return polymerizate(file, 40)
}

package main

import (
	"fmt"
	"unicode"

	"github.com/jdrst/adventofgo/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	lines := file.AsLines()
	exits := map[string][]string{}
	visits := map[string]int{}

	for _, l := range lines {
		caves := l.SubSplitWith("-")
		exits[string(caves[0])] = append(exits[string(caves[0])], string(caves[1]))
		exits[string(caves[1])] = append(exits[string(caves[1])], string(caves[0]))

	}
	return allPossiblePathsVisitOnce(exits, "start", visits)
}

func allPossiblePathsVisitOnce(exits map[string][]string, currentCave string, visits map[string]int) int {
	visits[currentCave] += 1
	sum := 0
	for _, c := range exits[currentCave] {
		if c == "end" {
			sum++
			continue
		}
		if !unicode.IsLower(rune(c[0])) || visits[c] == 0 {
			nextVisits := visits
			if unicode.IsLower(rune(c[0])) {
				nextVisits = copyMap(visits)
			}
			sum += allPossiblePathsVisitOnce(exits, c, nextVisits)
		}
	}
	return sum
}

func partTwo(file util.File) int {
	lines := file.AsLines()
	exits := map[string][]string{}
	visits := map[string]int{}

	for _, l := range lines {
		caves := l.SubSplitWith("-")
		exits[string(caves[0])] = append(exits[string(caves[0])], string(caves[1]))
		exits[string(caves[1])] = append(exits[string(caves[1])], string(caves[0]))
	}
	return allPossiblePathsOneDoubleVisit(exits, "start", visits, false)
}

func allPossiblePathsOneDoubleVisit(exits map[string][]string, currentCave string, visits map[string]int, hasVisitedTwice bool) int {
	if visits[currentCave] > 0 && unicode.IsLower(rune(currentCave[0])) {
		hasVisitedTwice = true
	}
	visits[currentCave] += 1
	sum := 0
	for _, c := range exits[currentCave] {
		if c == "end" {
			sum++
			continue
		}
		if c != "start" && (!unicode.IsLower(rune(c[0])) || visits[c] < 1 || !hasVisitedTwice) {
			nextVisits := visits
			if unicode.IsLower(rune(c[0])) {
				nextVisits = copyMap(visits)
			}
			sum += allPossiblePathsOneDoubleVisit(exits, c, nextVisits, hasVisitedTwice)
		}
	}
	return sum
}

func copyMap(visits map[string]int) map[string]int {
	res := make(map[string]int, len(visits))
	for k, v := range visits {
		res[k] = v
	}
	return res
}

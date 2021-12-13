package main

import (
	"github.com/jdrst/adventofgo/util"
)

func partOneWithUniquePathArray(file util.File) int {
	lines := file.AsLines()
	exits := map[string][]string{}

	for _, l := range lines {
		caves := l.SubSplitWith("-")
		exits[string(caves[0])] = append(exits[string(caves[0])], string(caves[1]))
		exits[string(caves[1])] = append(exits[string(caves[1])], string(caves[0]))
	}

	return allPossiblePathsVisitOnceWithUniquePathArray("start", exits, []string{"start"})
}

func allPossiblePathsVisitOnceWithUniquePathArray(current string, exits map[string][]string, path []string) int {
	sum := 0
loop:
	for _, next := range exits[current] {
		nextPath := path

		if next == "end" {
			sum++
			continue loop
		}

		if next[0] >= 'a' {
			if hasAlreadyVisited(path, next) {
				continue loop
			}
			nextPath = append(path, next)
		}

		sum += allPossiblePathsVisitOnceWithUniquePathArray(next, exits, nextPath)
	}
	return sum
}

func partTwoWithUniquePathArray(file util.File) int {
	lines := file.AsLines()
	exits := map[string][]string{}

	for _, l := range lines {
		caves := l.SubSplitWith("-")
		exits[string(caves[0])] = append(exits[string(caves[0])], string(caves[1]))
		exits[string(caves[1])] = append(exits[string(caves[1])], string(caves[0]))
	}

	return allPossiblePathsVisitTwiceWithUniquePathArray("start", exits, false, []string{"start"})
}

func allPossiblePathsVisitTwiceWithUniquePathArray(current string, exits map[string][]string, hasVisitedTwice bool, path []string) int {
	sum := 0
loop:
	for _, next := range exits[current] {
		nextPath := path
		alreadyVisited := false

		if next == "end" {
			sum++
			continue loop
		}

		if next == "start" {
			continue loop
		}

		if next[0] >= 'a' {
			alreadyVisited = hasAlreadyVisited(path, next)
			if alreadyVisited && hasVisitedTwice {
				continue loop
			}
			nextPath = append(path, next)
		}

		sum += allPossiblePathsVisitTwiceWithUniquePathArray(next, exits, hasVisitedTwice || alreadyVisited, nextPath)
	}
	return sum
}

func hasAlreadyVisited(path []string, next string) bool {
	for _, c := range path {
		if next == c {
			return true
		}
	}
	return false
}

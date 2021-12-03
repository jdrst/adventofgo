package main

import (
	"fmt"
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
	countOfOnes := make([]int, 12)
	for _, l := range lines {
		for i, c := range l {
			if c == '1' {
				countOfOnes[i]++
			}
		}
	}
	gma := strings.Builder{}
	eps := strings.Builder{}
	gammaRate := 0
	epsilonRate := 0
	for _, n := range countOfOnes {
		if n > len(lines)/2 {
			gma.WriteString(fmt.Sprintf("%v", 1))
			eps.WriteString(fmt.Sprintf("%v", 0))
			gammaRate = gammaRate<<1 | 1
			epsilonRate = epsilonRate << 1
		} else {
			gma.WriteString(fmt.Sprintf("%v", 0))
			eps.WriteString(fmt.Sprintf("%v", 1))
			gammaRate = gammaRate << 1
			epsilonRate = epsilonRate<<1 | 1

		}
	}
	fmt.Println(gma.String())
	fmt.Println(eps.String())
	return gammaRate * epsilonRate
}

func partTwo(file util.File) int {
	lines := file.AsLines()
	lines2 := file.AsLines()
	for i := 0; i < len(lines[0]); i++ {
		mostCommon := getMostCommonBitOnPos(i, lines)
		for j := 0; j < len(lines) && len(lines) > 1; j++ {
			if string(lines[j][i]) == "1" == mostCommon {
				continue
			}
			lines[j] = lines[len(lines)-1]
			lines = lines[:len(lines)-1]
			j--
		}
		leastCommon := getLeastCommonBitOnPos(i, lines2)
		for j := 0; j < len(lines2) && len(lines2) > 1; j++ {
			if string(lines2[j][i]) == "1" == leastCommon {
				continue
			}
			lines2[j] = lines2[len(lines2)-1]
			lines2 = lines2[:len(lines2)-1]
			j--
		}
	}
	first, err := strconv.ParseInt(string(lines[0]), 2, 64)
	util.Handle(err)
	second, err := strconv.ParseInt(string(lines2[0]), 2, 64)
	util.Handle(err)
	return int(first * second)
}

func getMostCommonBitOnPos(pos int, lines util.Lines) bool {
	cnt := 0
	for _, l := range lines {
		if l[pos] == '1' {
			cnt++
		}
	}
	return float64(cnt) >= float64(len(lines))/2
}

func getLeastCommonBitOnPos(pos int, lines util.Lines) bool {
	cnt := 0
	for _, l := range lines {
		if l[pos] == '1' {
			cnt++
		}
	}
	return float64(cnt) < float64(len(lines))/2
}

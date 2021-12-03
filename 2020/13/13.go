package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"math"
	"sort"
	"strconv"
	"strings"
	"github.com/deanveloper/modmath/v1/bigmod"
)

func parseInput() []byte {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return input
}

var newLine = "\r\n"

func prepareInput(input []byte) (int, []string) {
	lines := strings.Split(strings.TrimSpace(string(input)), newLine)
	departure, _ := strconv.Atoi(lines[0])
	buslines := strings.Split(lines[1], ",")

	return departure, buslines
}

func main() {
	departure, buslines := prepareInput(parseInput())

	offsets := makeOffsets(buslines)
	bestWaitTime, bestBusline := math.MaxInt64, 0
	for _, busline := range offsets {
		waitTime := busline - departure%busline
		if(waitTime < bestWaitTime) {
			bestWaitTime = waitTime
			bestBusline = busline
		}
	}
	fmt.Println(bestWaitTime*bestBusline)

	crtMat := makeCrtMat(offsets)
	fmt.Println(bigmod.SolveCrtMany(crtMat))
}

func makeOffsets(buslines []string) map[int]int {
	offsets := make(map[int]int)
	for offset, departure := range buslines {
		if departure != "x" {
			offsets[offset], _ = strconv.Atoi(departure)
		}
	}
	return offsets
}

func makeCrtMat(input map[int]int) []bigmod.CrtEntry {
	keys := make([]int, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	crtMat := make([]bigmod.CrtEntry, len(input))

	for i, k := range keys {
		crtMat[i] = bigmod.CrtEntry{A: big.NewInt(int64(input[k] - k)), N: big.NewInt(int64(input[k]))}
	}
	return crtMat
}

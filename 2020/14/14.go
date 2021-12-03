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

func prepareInput(input []byte) []string {
	lines := strings.Split(strings.TrimSpace(string(input)), newLine)

	return lines
}

func main() {
	lines := prepareInput(parseInput())

	fmt.Println(part1(lines))

	fmt.Println(part2(lines))
}

func part1(lines []string) uint64 {
	cache := make(map[uint64]uint64)
	var currentAndMask uint64
	var currentOrMask uint64
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			mask := line[7:]
			currentAndMask = parseUint(strings.ReplaceAll(mask, "X", "1"), 2, 36)
			currentOrMask = parseUint(strings.ReplaceAll(mask, "X", "0"), 2, 36)
			continue
		}
		if strings.HasPrefix(line, "mem") {
			var address, value uint64
			fmt.Sscanf(line, "mem[%d] = %d", &address, &value)
			cache[address] = value&currentAndMask | currentOrMask
		}
	}
	var sum uint64 = 0
	for _, value := range cache {
		sum += value
	}
	return sum
}

func part2(lines []string) uint64 {
	cache := make(map[uint64]uint64)
	currentMask := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			currentMask = line[7:]
		}
		if strings.HasPrefix(line, "mem") {
			var address, value uint64
			fmt.Sscanf(line, "mem[%d] = %d", &address, &value)
			permutatedAddresses := permutate(address, currentMask)
			for _, addr := range permutatedAddresses {
				cache[addr] = value
			}
		}
	}
	var sum uint64 = 0
	for _, value := range cache {
		sum += value
	}
	return sum
}

func permutate(address uint64, mask string) []uint64 {
	resAsStr := []string{""}
	appliedMask := ""
	bitAddress := strconv.FormatUint(address, 2)
	for i, j := len(mask)-1, len(bitAddress)-1; i >= 0; i, j = i-1, j-1 {
		if mask[i] == '0' && j >= 0 {
			appliedMask += string(bitAddress[j])
		} else {
			appliedMask += string(mask[i])
		}
	}
	for _, c := range appliedMask {
		currentLength := len(resAsStr)
		switch c {
		case 'X':
			resAsStr = append(resAsStr, resAsStr...)
			for i := 0; i < currentLength; i++ {
				resAsStr[i] = resAsStr[i] + "0"
			}
			for i := currentLength; i < len(resAsStr); i++ {
				resAsStr[i] = resAsStr[i] + "1"
			}
		default:
			for i := 0; i < len(resAsStr); i++ {
				resAsStr[i] = resAsStr[i] + string(c)
			}
		}
	}
	result := make([]uint64, len(resAsStr))
	for i, s := range resAsStr {
		result[i] = parseUint(s, 2, 0)
	}
	return result
}

func parseUint(s string, base, bitSize int) uint64 {
	result, err := strconv.ParseUint(s, base, bitSize)
	if err != nil {
		panic(err)
	}
	return result
}

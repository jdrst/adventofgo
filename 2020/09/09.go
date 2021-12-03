package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\r\n")
	ints := make([]int, len(lines))
	for i, num := range lines {
		asint, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		ints[i] = asint
	}

	findWeakNumber := func(ints []int) int {

	loop:
		for i := 25; i < len(ints); i++ {
			for j := i - 25; j < i; j++ {
				for k := j + 1; k < i; k++ {
					if ints[j]+ints[k] == ints[i] {
						continue loop
					}
				}
			}
			return ints[i]
		}
		return -1
	}

	findWeakness := func(ints []int, weakNumber int) int {
		for i := range ints {
			sum := 0
			j := 0
			for sum < weakNumber {
				sum += ints[i+j]
				if sum == weakNumber {
					sort.Ints(ints[i : i+j+1])
					return ints[i] + ints[i+j]
				}
				j++
			}

		}
		return -1
	}

	weakNumber := findWeakNumber(ints)
	fmt.Println(weakNumber)
	fmt.Println(findWeakness(ints, weakNumber))
}

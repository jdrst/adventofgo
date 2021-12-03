package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	handle(err)

	lines := strings.Split(string(input), "\r\n")
	for i, l := range lines {
		num1, err := strconv.Atoi(l)
		handle(err)
		for j, l2 := range lines[i+1:] {
			num2, err := strconv.Atoi(l2)
			handle(err)
			if num1+num2 == 2020 {
				fmt.Printf("Result 1: %v\n", num1*num2)
			}
			for _, l3 := range lines[i+j+1:] {
				num3, err := strconv.Atoi(l3)
				handle(err)
				if num1+num2+num3 == 2020 {
					fmt.Printf("Result 2: %v\n", num1*num2*num3)
				}
			}

		}
	}
}

func handle(e error) {
	if e != nil {
		panic(e)
	}
}

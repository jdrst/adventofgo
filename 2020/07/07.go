package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	bags := make(map[string]map[string]int)

	lines := strings.Split(string(input), "\r\n")
	for _, currentBag := range lines {
		f := strings.Fields(currentBag)
		color := strings.Join(f[0:2], " ")
		bagsWithAmounts := make(map[string]int)
		if !strings.Contains(currentBag, "no") {
			for i := 3; i < len(f); i++ {
				if strings.Contains(f[i], "bag") {
					amount, _ := strconv.Atoi(f[i-3])
					bagsWithAmounts[strings.Join(f[i-2:i], " ")] = amount
				}
			}
		}
		bags[color] = bagsWithAmounts
	}

	bagsRequiredInShinyGold := bagsRequiredIn("shiny gold", bags)
	bagsContainingShinyGold := make(map[string]int)
	bagsContaining("shiny gold", bags, bagsContainingShinyGold)
	fmt.Println(len(bagsContainingShinyGold))
	fmt.Println(bagsRequiredInShinyGold)
}

func bagsContaining(bagColor string, bags map[string]map[string]int, bagsHoldingBagColor map[string]int) {
	for color, b := range bags {
		if b[bagColor] > 0 {
			if bagsHoldingBagColor[color] <= 0 {
				bagsHoldingBagColor[color] = 1
				delete(bags, color)
				bagsContaining(color, bags, bagsHoldingBagColor)
			}
		}
	}
}

func bagsRequiredIn(startingBagColor string, bags map[string]map[string]int) (count int) {
	if len(bags[startingBagColor]) > 0 {
		for color, amount := range bags[startingBagColor] {
			count += amount * (bagsRequiredIn(color, bags) + 1)
		}
	}
	return
}

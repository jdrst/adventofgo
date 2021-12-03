package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	handle(err)
	validPasswords, validPasswordsNewPolicy := 0, 0

	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\r\n") {
		var password string
		var firstNum, secondNum int
		var char byte
		fmt.Sscanf(line, "%v-%v %c: %v", &firstNum, &secondNum, &char, &password)
		if count := strings.Count(password, string(char)); count >= firstNum && count <= secondNum {
			validPasswords++
		}
		if (password[firstNum-1] == char) != (password[secondNum-1] == char) {
			validPasswordsNewPolicy++
		}
	}

	fmt.Println(validPasswords)
	fmt.Println(validPasswordsNewPolicy)
}

func handle(e error) {
	if e != nil {
		panic(e)
	}
}

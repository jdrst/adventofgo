package main

import (
	"os"
	"testing"

	"github.com/jdrst/adventofgo/util"
)

func BenchmarkDay13(b *testing.B) {
	sout := os.Stdout
	os.Stdout = nil
	input := util.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		day13(input)
	}
	os.Stdout = sout
}

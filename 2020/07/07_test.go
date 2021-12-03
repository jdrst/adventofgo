package main

import "testing"

func Benchmark07(b *testing.B) {
	for n := 0; n < b.N; n++ {
		main()
	}
}

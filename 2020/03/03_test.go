package main

import "testing"

func Benchmark03(b *testing.B) {
	for n := 0; n < b.N; n++ {
		main()
	}
}

package main

import "testing"

var base = 10
var pow = 60

func BenchmarkPowerN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PowerN(base, pow)
	}
}

func BenchmarkPowerN2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PowerN2(base, pow)
	}
}

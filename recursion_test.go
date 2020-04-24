package main

import "testing"

var base = 10
var pow = 60
var big_num = 100000

func BenchmarkPowerN(b *testing.B) {
	b.Skip()
	for i := 0; i < b.N; i++ {
		PowerN(base, pow)
	}
}

func BenchmarkPowerN2(b *testing.B) {
	b.Skip()
	for i := 0; i < b.N; i++ {
		PowerN2(base, pow)
	}
}

func BenchmarkSlowAdd(b *testing.B) {
	b.Skip()
	for i := 0; i < b.N; i++ {
		SlowAdd(big_num, big_num+1)
	}
}

func BenchmarkSlowAdd1(b *testing.B) {
	b.Skip()
	for i := 0; i < b.N; i++ {
		SlowAdd1(big_num, big_num+1)
	}
}

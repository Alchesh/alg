package main

import "testing"

//var MaxCount = 100
var MaxCount = 1000

var RandomRange = 100

func BenchmarkSeqSearchDup(b *testing.B) {
	m := MyList{generate(1, MaxCount, RandomRange)}
	for i := 0; i < b.N; i++ {
		m.SeqSearchDup()
	}
}

func BenchmarkSortSearchDup(b *testing.B) {
	m := MyList{generate(1, MaxCount, RandomRange)}
	for i := 0; i < b.N; i++ {
		m.SortSearchDup()
	}
}

func BenchmarkHashSearchDup(b *testing.B) {
	m := MyList{generate(1, MaxCount, RandomRange)}
	for i := 0; i < b.N; i++ {
		m.HashSearchDup()
	}
}

func BenchmarkCountSearchDup(b *testing.B) {
	m := MyList{generate(1, MaxCount, RandomRange)}
	for i := 0; i < b.N; i++ {
		m.CountSearchDup()
	}
}

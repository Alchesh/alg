package main

import "testing"

var MaxCount = 1000

func BenchmarkSeqSearchDup(b *testing.B) {
	m := MyList{generate(1, MaxCount, 100)}
	for i := 0; i < b.N; i++ {
		m.SeqSearchDup()
	}
}

func BenchmarkSortSearchDup(b *testing.B) {
	m := MyList{generate(1, MaxCount, 100)}
	for i := 0; i < b.N; i++ {
		m.SortSearchDup()
	}
}

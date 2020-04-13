package main

import "testing"

func Test_SeqSearch(t *testing.T) {

	t.Parallel()

	m := MyList{[]int{1, 2, 8, 7, 5, 9, 6, 3, 4}}

	res := m.SeqSearch(5)

	if res != 8 {
		t.Error("SeqSearch error", 8, "!=", res)
	}

}

func Test_BinSearch(t *testing.T) {

	t.Parallel()

	m := MyList{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}}

	res := m.BinSearch(9)

	if res != 8 {
		t.Error("BinSearch error", 8, "!=", res)
	}

	res = m.BinSearch(1)

	if res != 0 {
		t.Error("BinSearch error", 0, "!=", res)
	}

	if m.BinSearch(0) != -1 && m.BinSearch(10) != -1 {
		t.Error("BinSearch board error", m.BinSearch(0), "!=", m.BinSearch(10), "!=", -1)
	}

}
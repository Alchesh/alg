package main

import (
	"reflect"
	"testing"
)

func Test_SeqSearch(t *testing.T) {
	t.Skip()
	t.Parallel()

	m := MyList{[]int{1, 2, 8, 7, 5, 9, 6, 3, 4}}

	res := m.SeqSearch(5)

	if res != 8 {
		t.Error("SeqSearch error", 8, "!=", res)
	}

}

func Test_BinSearch(t *testing.T) {
	t.Skip()
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

func Test_Reverse(t *testing.T) {
	t.Skip()
	t.Parallel()

	ar := []int{1, 2, 8, 7, 5, 9, 6, 3, 4}
	m := MyList{ar}

	m.Reverse(0, len(ar)-1)

	if !reflect.DeepEqual(ar, m.Get()) {
		t.Error("Reverse error\n", ar, "\n!=\n", m.Get())
	}

}

func Test_Rotate(t *testing.T) {
	//t.Skip()
	t.Parallel()

	m := MyList{[]int{1, 2, 3, 4, 5, 6}}
	m.Rotate(2)

	if !reflect.DeepEqual([]int{3, 4, 5, 6, 1, 2}, m.Get()) {
		t.Error("Rotate left error\n", m.Get())
	}

	m = MyList{[]int{1, 2, 3, 4, 5, 6}}
	m.Rotate(-2)

	if !reflect.DeepEqual([]int{5, 6, 1, 2, 3, 4}, m.Get()) {
		t.Error("Rotate right error\n", m.Get())
	}

}

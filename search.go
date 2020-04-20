package main

import (
	//	"fmt"
	"sort"
)

// SeqSearchDup - sequence duplicates search
func (m *MyList) SeqSearchDup() {

	l := len((*m).mas)

	for i, v := range (*m).mas {
		for _, w := range (*m).mas[i+1 : l] {
			if v == w {
				//fmt.Println("Find duples", w)
				continue
			}
		}
	}
}

// SeqSearchDup - sort & search
func (m *MyList) SortSearchDup() {
	l := len((*m).mas)
	sort.Ints((*m).mas)

	for i := 1; i < l; i++ {
		if (*m).mas[i] == (*m).mas[i-1] {
			//fmt.Println("Find duples", (*m).mas[i-1])
			continue
		}
	}
}

// SeqSearchDup - hash duplicates search
func (m *MyList) HashSearchDup() {

	set := make(map[int]bool)
	l := len((*m).mas)

	for i := 0; i < l; i++ {
		if _, ok := set[(*m).mas[i]]; ok {
			//fmt.Println("Find duples", (*m).mas[i])
			continue
		} else {
			set[(*m).mas[i]] = true
		}

	}
}

// SeqSearchDup - duplicates search by count
func (m *MyList) CountSearchDup() {

	l := len((*m).mas)
	count := make([]int, 100)

	for i := 0; i < l; i++ {
		if count[(*m).mas[i]] > 0 {
			//fmt.Println("Find duples", (*m).mas[i])
			continue
		} else {
			count[(*m).mas[i]]++
		}
	}
}

// SeqCalcMax - sequence calculate maximum number of times
func (m *MyList) SeqCalcMax() (int, int) {

	var el int
	MaxCount := 0
	count := 1

	for i, v := range (*m).mas {
		for _, w := range (*m).mas[i+1:] {
			if v == w {
				count++
			}
		}
		if count > MaxCount {
			MaxCount = count
			el = v
			count = 1
		}
	}

	return MaxCount, el
}

// SortCalcMax - sort & calculate maximum number of times
func (m *MyList) SortCalcMax() (int, int) {
	l := len((*m).mas)
	sort.Ints((*m).mas)

	el := 0
	MaxCount := 0
	count := 1
	cur_el := (*m).mas[0]

	for i := 1; i < l; i++ {
		if (*m).mas[i] == cur_el {
			count++
		} else {
			count = 1
			cur_el = (*m).mas[i]
		}
		if count > MaxCount {
			MaxCount = count
			el = cur_el
		}
	}

	return MaxCount, el
}

// SeqSearchDup - duplicates search by count
func (m *MyList) CountCalcMax() (int, int) {

	el := 0
	MaxCount := 0
	l := len((*m).mas)
	count := make([]int, 100)

	for i := 0; i < l; i++ {
		count[(*m).mas[i]]++
	}

	for i := 0; i < len(count); i++ {
		if count[i] > MaxCount {
			MaxCount = count[i]
			el = i
		}
	}

	return MaxCount, el
}

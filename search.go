package main

import (
	//	"fmt"
	"sort"
)

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

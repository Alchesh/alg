package main

import "fmt"

// MyList - simple array
type MyList struct {
	mas []int
}

// Sum - summarize all elements
func (m *MyList) Sum() int {
	var s int

	for _, v := range (*m).mas {
		s += v
	}

	return s
}

// SeqSearch - Sequantional search
func (m *MyList) SeqSearch(val int) int {

	res := -1

	for i, v := range (*m).mas {
		if v == val {
			res = i
			break
		}
	}

	return res
}

// BinSearch - Binary search
func (m *MyList) BinSearch(val int) int {

	res := -1

	l := 0
	r := len((*m).mas)

	for l < r {
		// c := l + (r - l) / 2 // for addition overflow
		c := (l + r) / 2
		if (*m).mas[c] == val {
			res = c
			break
		} else {
			if (*m).mas[c] > val {
				r = c
			} else {
				l = c + 1
			}
		}
	}

	return res
}

// Rotate - Rotate array
// [1, 2, 3, 4, 5, 6] -> by 2 -> positions to [3, 4, 5, 6, 1, 2]
func (m *MyList) Rotate(val int) {

	//for _, _ := range (*m).mas {

	//}

}

func main() {

	fmt.Println("All Done!")
}

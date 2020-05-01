package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

// Get - get array
func (m *MyList) Get() []int {
	return (*m).mas
}

// Reverse - reverse array
func (m *MyList) Reverse(i, j int) {

	for i < j {
		(*m).mas[i], (*m).mas[j] = (*m).mas[j], (*m).mas[i]
		i++
		j--
	}

}

// Rotate - Rotate array
// [1, 2, 3, 4, 5, 6] -> by 2 -> positions to [3, 4, 5, 6, 1, 2]
func (m *MyList) Rotate(val int) {

	n := len((*m).mas) - 1
	if val > 0 {
		m.Reverse(0, val-1) // 2 1 3 4 5 6 (to left)
		m.Reverse(val, n)   // 2 1 6 5 4 3
	} else {
		val = -val
		m.Reverse(n-val+1, n) // 1 2 3 4 6 5 (to right)
		m.Reverse(0, n-val)   // 4 3 2 1 5 6
	}
	m.Reverse(0, n) // complete
}

func generate(t, size, max_val int) []int {

	ar := make([]int, size)

	switch t {
	case 0: //sorted
		for i := 0; i < size; i++ {
			ar[i] = i
		}
	case 1: //not sorted
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < size; i++ {
			ar[i] = rand.Intn(max_val)
		}
	}

	return ar
}

func main() {

	m := MyList{generate(1, 50, 10)}
	// m := MyList{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}}
	// m := MyList{[]int{1, 2, 3, 4, 5, 6, 9, 8, 9}}
	// m := MyList{[]int{1, 2, 3, 4, 5, 6, 9, 8, 9}}
	//m := MyList{[]int{1, 2, 3, 4, 5, 6, 9, 8, 9, 1, 2, 3, 4, 5, 6, 9, 8, 0, 9}}

	f := more
	//f := less
	//m.BubbleSort1(f)
	m.InsertionSort(f)
	fmt.Println(m.mas)

	fmt.Println("\nAll Done!")
}

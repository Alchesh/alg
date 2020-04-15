package main

func fakk(val int) int {

	if val < 2 {
		return 1
	}

	return val * fakk(val-1)
}

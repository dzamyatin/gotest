package lib

import rand "math/rand"

func CreateRandList(count int, border int) []int {
	res := make([]int, count)

	for k, _ := range res {
		res[k] = rand.Intn(border)
	}

	return res
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(Partitions(5))
}

func Partitions(n int) int {

	m := make(map[string]struct{})
	elem := make([]int, n)
	res := 0

	fil(elem, 0, &res, m)
	//fmt.Println(elem)

	return len(m)
}

func fil(elem []int, ps int, res *int, m map[string]struct{}) {
	nps := ps + 1
	for i := 0; i <= len(elem); i++ {
		elem[ps] = i
		if nps < len(elem) {
			fil(elem, nps, res, m)
		}
		//fmt.Println(elem)

		s := 0
		for _, v := range elem {
			s += v
		}
		if s == len(elem) {
			*res = *res + 1
			e := make([]int, len(elem))
			copy(e, elem)
			sort.Ints(e)
			m[fmt.Sprint(e)] = struct{}{}
		}
	}
}

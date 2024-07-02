package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(Partitions(5))
	//fmt.Println(Partitions(25))
}

func Partitions(n int) int {

	m := make(map[string]struct{})
	elem := make([]int, n)

	fil(elem, 0, m)

	return len(m)
}

func fil(elem []int, ps int, m map[string]struct{}) {
	nps := ps + 1
	for i := len(elem); i >= 0; i-- {
		elem[ps] = i

		//fmt.Println(elem)

		s := sum(elem)

		if s < len(elem) {
			if nps < len(elem) {
				fil(elem, nps, m)
			}
		}

		if s == len(elem) {
			store(elem, m)
		}
	}
}

func store(elem []int, m map[string]struct{}) {
	e := make([]int, len(elem))
	copy(e, elem)
	sort.Ints(e)
	m[fmt.Sprint(e)] = struct{}{}
}

func sum(elem []int) int {
	s := 0
	for _, v := range elem {
		s += v
	}
	return s
}

//func fil(elem []int, ps int, res *int, m map[string]struct{}) {
//	nps := ps + 1
//	for i := 0; i <= len(elem); i++ {
//		elem[ps] = i
//
//		s := 0
//		for _, v := range elem {
//			s += v
//		}
//
//		if s >= len(elem) {
//			break
//		}
//
//		fmt.Println(elem)
//
//		if s == len(elem) {
//			*res = *res + 1
//			e := make([]int, len(elem))
//			copy(e, elem)
//			sort.Ints(e)
//			m[fmt.Sprint(e)] = struct{}{}
//		}
//
//		if nps < len(elem) {
//			fil(elem, nps, res, m)
//		}
//	}
//}

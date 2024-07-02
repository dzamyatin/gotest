package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	//fmt.Println(Partitions(1))
	//fmt.Println(Partitions(5))
	fmt.Println(Partitions(10))
	//fmt.Println(Partitions(25))
}

// [5], [4,1], [3,2], [3,1,1], [2,2,1], [2,1,1,1], [1,1,1,1,1]
// [1,1,1,1,1] 0/1 => [1,1,1,1,1] //
// [1,1,1,1,1] 0/2 => [2,2,1] // 1/2 [2,3]
// [1,1,1,1,1] 0/3 => [3,1,1] // 1/2 [3,2]
// [1,1,1,1,1] 0/4 => [4,1]
// [1,1,1,1,1] 0/5 => [5]
func Partitions(n int) int {

	m := make(map[string]struct{})

	r := make([]int, n)
	for i := range r {
		r[i] = 1
	}

	//fmt.Println(
	//	conv(r, 1, 0),
	//	conv(r, 2, 0),
	//	conv(r, 3, 0),
	//	conv(r, 4, 0),
	//	conv(r, 5, 0),
	//
	//	conv(r, 1, 1),
	//	conv(r, 2, 1),
	//	conv(r, 3, 1),
	//	conv(r, 4, 1),
	//	conv(r, 5, 1),
	//)

	rec(r, m, 0, true)

	log.Println(m)

	return len(m)
}

func rec(elem []int, m map[string]struct{}, g int, flag bool) {
	for i := 1; i <= len(elem); i++ {
		r := conv(elem, i, g)
		store(r, m)

		if len(r) != len(elem) || flag {
			for j := 0; j <= len(elem); j++ {
				rec(r, m, j, false)
			}
		}
	}
}

func conv(elem []int, group int, gap int) []int {
	var r []int
	j := 0

	var buf []int
	for i := 0; i < len(elem); i++ {

		if i >= gap {
			j++
		}
		buf = append(buf, elem[i])
		if j%group == 0 {
			r = append(r, sum(buf))
			buf = []int{}
			continue
		}
	}

	if sum(buf) > 0 {
		r = append(r, sum(buf))
	}

	return r
}

func sum(ele []int) int {
	var s int
	for i := 0; i < len(ele); i++ {
		s += ele[i]
	}
	return s
}

func store(r []int, m map[string]struct{}) {
	sort.Ints(r)
	m[fmt.Sprint(r)] = struct{}{}
}

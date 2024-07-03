package main

import (
	"fmt"
	"sort"
)

var stat int

func main() {
	//fmt.Println(Partitions(1))
	//fmt.Println(Partitions(5))
	//fmt.Println(Partitions(10))
	fmt.Println(Partitions(25))
	//
	fmt.Println("ST", stat)
}

// [5], [4,1], [3,2], [3,1,1], [2,2,1], [2,1,1,1], [1,1,1,1,1]
//
// [1,1,1,1,1] => 1-5
// [1,1,1,1,1] => 2-2,
// [1,1,1,1,1] => 3-1,
// [1,1,1,1,1] => 4-1,
// [1,1,1,1,1] => 5-1,

var cyclePath = make(map[string]int)

func Partitions(n int) int {

	var sum int
	var storage = make(map[string]struct{})

	m := make([]int, n)

	for i := 0; i < n; i++ {
		m[i] = i + 1
	}

	fmt.Println(m)

	rec(m, 0, []int{}, &sum, storage)

	return sum
}

func rec(m []int, cur int, path []int, sum *int, storage map[string]struct{}) {
	stat++
	for _, v := range m {
		newPath := make([]int, len(path))

		copy(newPath, path)
		newPath = append(newPath, v)

		if overfull(newPath) {
			continue
		}

		if cur+v == len(m) {
			store(newPath, sum, storage)
			continue
		}
		if cur+v < len(m) {
			rec(m, cur+v, newPath, sum, storage)
		}
	}
}

func overfull(r []int) bool {
	fmt.Println(r)

	sort.Ints(r)
	k := fmt.Sprint(r)

	v := cyclePath[k]

	cyclePath[k]++

	return v > 0
}

func store(r []int, sum *int, storage map[string]struct{}) {
	sort.Ints(r)
	k := fmt.Sprint(r)
	if _, ok := storage[k]; ok {
		return
	}
	//fmt.Println(r)
	storage[k] = struct{}{}
	*sum++
}

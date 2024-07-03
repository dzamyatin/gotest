package main

import (
	"fmt"
	"sort"
)

//var stat int

func main() {
	//fmt.Println(Partitions(1))
	//fmt.Println(Partitions(5))
	fmt.Println(Partitions(10))
	//fmt.Println(Partitions(25))
	//fmt.Println(Partitions(100))
	//
	//fmt.Println("ST", stat)
}

// [5], [4,1], [3,2], [3,1,1], [2,2,1], [2,1,1,1], [1,1,1,1,1]
//
// [1,1,1,1,1] => 1-5
// [1,1,1,1,1] => 2-2,
// [1,1,1,1,1] => 3-1,
// [1,1,1,1,1] => 4-1,
// [1,1,1,1,1] => 5-1,

func Partitions(n int) int {
	np := NewPart(n)
	return np.run()
}

type Part struct {
	sum     int
	storage map[string]struct{}
	cycle   map[string]struct{}
	number  int
}

func NewPart(number int) Part {
	return Part{
		storage: make(map[string]struct{}),
		cycle:   make(map[string]struct{}),
		number:  number,
	}
}

func (p *Part) run() int {
	m := make([]int, p.number)

	for i := 0; i < p.number; i++ {
		m[i] = i + 1
	}

	p.rec(m, 0, []int{})

	return p.sum
}

func (p *Part) rec(m []int, cur int, path []int) {
	for _, v := range m {
		newPath := make([]int, len(path))

		copy(newPath, path)
		newPath = append(newPath, v)

		if p.overfull(newPath) {
			continue
		}

		if cur+v == len(m) {
			p.store(newPath)
			continue
		}
		if cur+v < len(m) {
			p.rec(m, cur+v, newPath)
		}
	}
}

func (p *Part) store(r []int) {
	sort.Ints(r)
	k := fmt.Sprint(r)
	if _, ok := p.storage[k]; ok {
		return
	}
	//fmt.Println(r)
	p.storage[k] = struct{}{}
	p.sum++
}

func (p *Part) overfull(r []int) bool {
	sort.Ints(r)
	k := fmt.Sprint(r)

	if _, ok := p.cycle[k]; ok {
		return true
	}

	p.cycle[k] = struct{}{}

	return false
}

///---------------------------------------------

var cyclePath = make(map[string]int)

func Tst(n int) int {

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
	//stat++
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
	//fmt.Println(r)

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

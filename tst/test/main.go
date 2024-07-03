package main

import (
	"fmt"
	"sort"
)

//var stat int

func main() {
	//fmt.Println(Partitions(1))
	//fmt.Println(Partitions(5))
	//fmt.Println(Partitions(10))
	//fmt.Println(Partitions(25))
	fmt.Println(Partitions(40))
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

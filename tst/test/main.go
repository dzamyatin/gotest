package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sort"
	"time"
)

func main() {
	f, _ := os.Create("cpu.prof")
	err := pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatal("Write cpu profiling started error")
	}

	fs, _ := os.Create("heap.prof")
	err = pprof.WriteHeapProfile(fs)
	if err != nil {
		log.Fatal("Write heap profiling started error")
	}

	t := time.Now()
	//fmt.Println(Partitions(1))
	fmt.Println(Partitions(5))
	//fmt.Println(Partitions(10))
	//fmt.Println(Partitions(25))
	//fmt.Println(Partitions(40))
	//fmt.Println(Partitions(50))
	//fmt.Println(Partitions(100))

	fmt.Println("ST", time.Now().Sub(t))
	pprof.StopCPUProfile()
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
	sum    int
	cycle  map[string]struct{}
	number int
}

func NewPart(number int) Part {
	return Part{
		cycle:  make(map[string]struct{}),
		number: number,
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
			p.sum++
			continue
		}
		if cur+v < len(m) {
			p.rec(m, cur+v, newPath)
		}
	}
}

func (p *Part) overfull(r []int) bool {
	sort.Ints(r)

	var b = make([]byte, 0, 100)
	for _, v := range r {
		b = append(b, byte(v))
	}

	key := string(b)

	if _, ok := p.cycle[key]; ok {
		return true
	}

	p.cycle[key] = struct{}{}

	return false
}

//func (p *Part) overfull(r []int) bool {
//	sort.Ints(r)
//
//	b := strings.Builder{}
//	for _, v := range r {
//		b.WriteByte(byte(v))
//	}
//	k := b.String()
//
//	if _, ok := p.cycle[k]; ok {
//		return true
//	}
//
//	p.cycle[k] = struct{}{}
//
//	return false
//}

//func (p *Part) overfull(r []int) bool {
//	sort.Ints(r)
//	k := fmt.Sprint(r)
//
//	if _, ok := p.cycle[k]; ok {
//		return true
//	}
//
//	p.cycle[k] = struct{}{}
//
//	return false
//}

///---------------------------------------------

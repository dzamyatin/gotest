package main

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

//var stat int

func main() {
	t := time.Now()
	//fmt.Println(Partitions(1))
	fmt.Println(Partitions(5))
	fmt.Println(Partitions(10))
	fmt.Println(Partitions(25))
	fmt.Println(Partitions(40))
	//fmt.Println(Partitions(100))

	fmt.Println("ST", time.Now().Sub(t))
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
	sm     chan struct{}
	wg     sync.WaitGroup
	sum    atomic.Uint32
	cycle  sync.Map
	number int
}

func NewPart(number int) Part {
	return Part{
		sm:     make(chan struct{}, 10),
		cycle:  sync.Map{},
		number: number,
	}
}

func (p *Part) run() int {
	m := make([]int, p.number)

	for i := 0; i < p.number; i++ {
		m[i] = i + 1
	}

	p.rec(m, 0, []int{})
	p.wg.Wait()
	close(p.sm)
	return int(p.sum.Load())
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
			p.sum.Add(1)
			continue
		}
		if cur+v < len(m) {
			if cur == 0 {
				p.wg.Add(1)
				p.sm <- struct{}{}
				go func() {
					p.rec(m, cur+v, newPath)
					<-p.sm
					p.wg.Done()
				}()
			} else {
				p.rec(m, cur+v, newPath)
			}
		}
	}
}

func (p *Part) overfull(r []int) bool {
	sort.Ints(r)
	k := fmt.Sprint(r)

	_, loaded := p.cycle.LoadOrStore(k, struct{}{})

	return loaded
}

///---------------------------------------------

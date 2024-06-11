package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(index int, wgr *sync.WaitGroup) {
			fmt.Println(index)
			wgr.Done()
		}(i, &wg)
	}

	wg.Wait()
}

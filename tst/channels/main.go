package main

import (
	"app/tst/lib"
	"context"
	"fmt"
	"sync"
)

func main() {
	workers := 100

	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	chOut := make(chan []int)
	for i := 0; i < workers; i++ {
		inChan := make(chan []int)
		wg.Add(1)
		go func() {
			select {
			case <-ctx.Done():
				defer wg.Done()
				return
			case base := <-inChan:
				defer wg.Done()
				chOut <- lib.BTreeSortAsyncAdapter(base, lib.ASC)
				return
			}
		}()

		go func() {
			select {
			case <-ctx.Done():
				close(inChan)
				return
			default:
				inChan <- lib.CreateRandList(10000000, 10)
				close(inChan)
				return
			}
		}()
	}

	go func() {
		wg.Wait()
		close(chOut)
	}()

	i := 0
	for range chOut {
		i++

		if i > 1 {
			cancel()
		}

		fmt.Printf("Done: %v \n", i)
	}

}

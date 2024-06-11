package main

import (
	"app/tst/lib"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}
	s := lib.NewSemaphoreSync(5)

	tst := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			s.Lock()
			defer s.Release()

			mx.Lock()
			tst++
			fmt.Printf("Test %v\n\n", tst)
			mx.Unlock()

			time.Sleep(time.Second * 5)
		}()
	}

	wg.Wait()
}

//Task: context if done  == 0 else res

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	thingCtx, _ := context.WithTimeout(ctx, 2*time.Second)

	fmt.Println(
		doTheThing(thingCtx),
	)
}

func doTheThing(ctx context.Context) int {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Second * 5)
		ch <- 1
	}()

	select {
	case v := <-ch:
		return v
	case <-ctx.Done():
		return 0
	}
}

//Task: context if done  == 0 else res

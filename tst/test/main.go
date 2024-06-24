package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	_ "golang.org/x/sync/errgroup"
	"log"
	"time"
)

func main() {
	ctx, cancelFn := context.WithTimeout(context.Background(), 2*time.Second)
	group, _ := errgroup.WithContext(ctx)
	//group := errgroup.Group{}

	group.SetLimit(10)

	for i := 0; i < 100; i++ {
		group.Go(func() error {
			select {
			case <-ctx.Done():
				cancelFn()
				return nil
			default:
				log.Println(i)
				time.Sleep(5 * time.Second)
				return fmt.Errorf("asd")
			}
		})
	}

	err := group.Wait()

	if err != nil {
		log.Fatal(err)
	}
}

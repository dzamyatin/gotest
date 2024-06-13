package main

import "log"

func main() {

	ch := make(chan int, 1)

	ch <- 1

	close(ch)

	for v := range ch {
		log.Println(v)
	}

	log.Println("--")

	tst, ok := <-ch
	log.Println(tst, ok)
}

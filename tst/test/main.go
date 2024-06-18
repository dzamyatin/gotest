package main

import (
	"app/tst/lib"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	db := lib.NewKVDB()

	cache := lib.NewKVCache(db)
	//cache := lib.NewKVDBAsyncMap(db)

	var klist [100]string

	for i := 0; i < len(klist); i++ {
		klist[i] = "tst_" + strconv.Itoa(i)
	}

	for i := 0; i < 100; i++ {
		go func(i int) {
			cache.Set(getListEl(klist, i), strconv.Itoa(i))
		}(i)
	}

	strt := time.Now()
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func(ix int) {
			ix = getIx(ix, len(klist))
			cache.Get(getListEl(klist, i))
			wg.Done()
		}(i)
	}

	wg.Wait()

	log.Println("Met:")
	log.Println(lib.KvMetric.Load())
	log.Println("Met2:")
	log.Println(lib.KvMetric2.Load())
	end := time.Now().Sub(strt)
	fmt.Printf("Time: %s", end)
}

func getListEl(klist [100]string, i int) string {
	return klist[getIx(i, len(klist))]
}

func getIx(ix int, len int) int {
	if ix < len {
		return ix
	}

	for {
		ix = ix - len

		if ix < len {
			return ix
		}
	}
}

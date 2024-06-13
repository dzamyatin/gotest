package main

import (
	"app/tst/lib"
	"fmt"
	"time"
)

type SortVariant struct {
	title    string
	callback SortCallback
}

func main() {
	sorts := []SortVariant{
		{"SLOW", lib.SlowSort},
		{"MERGE", lib.QuickSort},
		{"Async MERGE", lib.QuickSortAsync}, //have better performance then MERGE on arrays large then 100000 elements
		{"BTree", lib.BTreeSortAdapter},
		{"Async BTree", lib.BTreeSortAsyncAdapter}, //have better performance in all cases
		{"Native sort", lib.NativeSortAdapter},
	}

	for _, v := range sorts {
		fmt.Printf("\n ------------------------------------------- \n Type: %v \n", v.title)
		t := time.Now()

		//testSort(v, 100000000, 2)
		testSort(v.callback, 20, 200)

		fmt.Printf("Time: %s \n-------------------------------------------\n", time.Now().Sub(t))
	}

}

type SortCallback func(base []int, order lib.SortOrder) []int

func testSort(sortCallback SortCallback, step int, elemNumber int) {
	size := 15

	points := make([]int, elemNumber)
	for i := 0; i < elemNumber; i++ {
		ce := i * step
		k := i

		test := lib.CreateRandList(ce, 1000)

		start := time.Now()

		sortCallback(test, lib.DESC)

		end := time.Now()

		elapsed := end.Sub(start).Nanoseconds()

		points[k] = int(elapsed)
	}

	tst := print(points, size)
	draw(print(tst, size), size)
}

func print(points []int, sizeLines int) []int {

	var maximum = 0
	for k, v := range points {
		if k == 0 {
			maximum = v
			continue
		}

		if v > maximum {
			maximum = v
		}
	}

	res := make([]int, len(points))
	for k, v := range points {
		res[k] = (v * sizeLines) / maximum
	}

	return res
}

func draw(points []int, size int) {
	for i := size; i >= 0; i-- {
		for _, v := range points {
			draw := "."
			if v == i {
				draw = "@"
			}
			fmt.Print(draw)
		}
		fmt.Print("\n")
	}
}

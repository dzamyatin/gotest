package main

import (
	"fmt"
	rand2 "math/rand"
	"time"
)

func main() {

	sorts := map[string]SortCallback{
		"SLOW":  slowSort,
		"MERGE": quickSort,
	}

	for k, v := range sorts {
		fmt.Printf("\n ------------------------------------------- \n Type: %v \n", k)
		t := time.Now()

		testSort(v, 1000, 200)

		fmt.Printf("Time: %s \n-------------------------------------------\n", time.Now().Sub(t))
	}

}

type SortCallback func(base []int, order SortOrder) []int

func testSort(sortCallback SortCallback, step int, elemNumber int) {
	size := 15

	points := make([]int, elemNumber)
	for i := 0; i < elemNumber; i++ {
		ce := i + step
		k := i

		test := createRandList(ce, 1000)

		start := time.Now()

		sortCallback(test, DESC)

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

func quickSort(base []int, order SortOrder) []int {
	// count = 3
	// 5, 3, 0
	// max = 5
	// min = 0
	// middle = (max - min) / count
	// 		  =  5  -  0 = 5 / 2 = 2.5
	// if v > middle => left else => right

	left, right := makeParts(base, order)

	if (len(left)+len(right) > 3) && len(left) != 0 && len(right) != 0 {
		return append(
			quickSort(left, order),
			quickSort(right, order)...,
		)
	}

	return slowSort(append(left, right...), order)
}

var metric = 0

func makeParts(list []int, order SortOrder) (left []int, right []int) {
	metric++
	var minimum int
	var maximum int

	for k, v := range list {
		if k == 0 {
			minimum = v
			maximum = v
		}

		if v < minimum {
			minimum = v
		}

		if v > maximum {
			maximum = v
		}
	}

	middle := (maximum - minimum) / 2
	for _, v := range list {

		var isBetter bool
		switch order {
		case ASC:
			isBetter = v > middle
		case DESC:
			isBetter = v < middle
		default:
			isBetter = v < middle
		}

		if isBetter {
			right = append(right, v)
			continue
		}
		left = append(left, v)
	}

	return
}

type SortOrder int8

const (
	ASC SortOrder = iota
	DESC
)

func slowSort(base []int, order SortOrder) []int {
	var containerValue int
	var containerKey int

	skipList := make(map[int]interface{})
	result := make([]int, len(base))

	for index, _ := range result {
		firstElement := true

		for k, v := range base {
			_, ok := skipList[k]

			if ok {
				continue
			}

			var isBetter bool
			switch order {
			case ASC:
				isBetter = v < containerValue
			case DESC:
				isBetter = v > containerValue
			default:
				isBetter = v < containerValue
			}

			if isBetter || firstElement {
				containerValue = v
				containerKey = k
			}

			firstElement = false
		}

		skipList[containerKey] = nil
		result[index] = containerValue
	}

	return result
}

func createRandList(count int, border int) []int {
	res := make([]int, count)

	for k, _ := range res {
		res[k] = rand2.Intn(border)
	}

	return res
}
package lib

import "sync"

func QuickSort(base []int, order SortOrder) []int {
	// count = 3
	// 5, 3, 0
	// max = 5
	// min = 0
	// middle = (max - min) / count
	// 		  =  5  -  0 = 5 / 2 = 2.5
	// if v > middle => left else => right

	left, right := makeParts(base, order)

	if (len(left)+len(right) > 2) && len(left) != 0 && len(right) != 0 {
		return append(
			QuickSort(left, order),
			QuickSort(right, order)...,
		)
	}

	return append(left, right...)
}

func QuickSortAsync(base []int, order SortOrder) []int {

	left, right := makeParts(base, order)

	if (len(left)+len(right) > 2) && len(left) != 0 && len(right) != 0 {

		wg := sync.WaitGroup{}

		wg.Add(2)

		var qSortA []int
		go func() {
			defer wg.Done()
			qSortA = QuickSortAsync(left, order)
		}()
		var qSortB []int
		go func() {
			defer wg.Done()
			qSortB = QuickSortAsync(right, order)
		}()
		wg.Wait()
		return append(
			qSortA,
			qSortB...,
		)
	}

	return append(left, right...)
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

	middle := minimum + ((maximum - minimum) / 2)
	for _, v := range list {

		var isBetter bool
		switch order {
		case ASC:
			isBetter = v > middle
		case DESC:
			isBetter = v <= middle
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

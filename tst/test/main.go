package main

import (
	"fmt"
	_ "golang.org/x/sync/errgroup"
	"log"
	"sort"
)

func main() {

	//tst := []int{
	//	1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
	//}
	//fmt.Println(tst[2:3])
	//return
	fmt.Println(
		SumOfIntervals(
			[][2]int{
				{1, 4},
				{3, 5},
				{7, 10},
			},
		),
		SumOfIntervals(
			[][2]int{
				{1, 4},
				{7, 10},
				{3, 5},
			},
		),

		SumOfIntervals(
			[][2]int{
				{1, 2},
				{2, 3},
				{3, 4},
			},
		),
	)
}

func SumOfIntervals(intervals [][2]int) int {
	var l []int

	for _, v := range intervals {
		l = append(l, v[:]...)
	}

	sort.Ints(l)

	for j := 0; j < len(l); j++ {

		currIntervalIndex := 0
		currIntervalFound := false

		for i := 0; i < len(intervals); i++ {
			if l[j] == intervals[i][0] {
				currIntervalIndex = i
				currIntervalFound = true
				break
			}
		}

		if !currIntervalFound {
			continue
		}

		if j == len(l)-1 {
			break
		}

		for i := 0; i < len(intervals); i++ {
			if currIntervalIndex == i {
				continue
			}

			//1,2
			//2,4
			if intervals[currIntervalIndex][0] <= intervals[i][0] &&
				intervals[currIntervalIndex][1] >= intervals[i][0] {
				r := [][2]int{}

				//left border
				if currIntervalIndex < i {
					if currIntervalIndex-1 > 0 {
						r = append(
							r,
							intervals[:currIntervalIndex-1]...,
						)
					}
				} else {
					if i-1 > 0 {
						r = append(
							r,
							intervals[:i-1]...,
						)
					}
				}

				//new
				r = append(
					r,
					[2]int{
						intervals[currIntervalIndex][0],
						intervals[i][1],
					},
				)

				//middle
				r = append(
					r,
					intervals[currIntervalIndex+1:i]...,
				)

				//right border
				if currIntervalIndex > i {
					if currIntervalIndex+1 <= len(intervals)-1 {
						r = append(
							r,
							intervals[currIntervalIndex+1:]...,
						)
					}
				} else {
					if i+1 <= len(intervals)-1 {
						r = append(
							r,
							intervals[i+1:]...,
						)
					}
				}

				intervals = r

				j = -1
				break
			}
		}
	}
	log.Println(intervals)
	res := 0
	return res
}

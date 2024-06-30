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

	//tst := 0
	tstI := 0
	//skip := 0

	for j := 0; j < len(l); j++ {
		//if l[j] <= skip {
		//	continue
		//}

		for i := 0; i < len(intervals); i++ {
			if l[j] == intervals[i][0] {
				//tst = intervals[i][1]
				tstI = i
				break
			}
		}

		if j == len(l)-1 {
			break
		}

		//if l[j+1] == tst {
		//	continue
		//}

		for i := 0; i < len(intervals); i++ {
			if l[j+1] == intervals[i][0] {
				//<<<
				r := append(
					[][2]int{},
					intervals[:tstI]...,
				)

				r = append(
					r,
					[2]int{
						intervals[tstI][0],
						intervals[i][1],
					},
				)

				//skip = intervals[i][1]

				if tstI+1 < i {
					r = append(
						r,
						intervals[tstI+1:i]...,
					)
				}

				intervals = append(
					r,
					intervals[i+1:]...,
				)

				//j = 0
				break
			}
		}
	}
	log.Println(intervals)
	res := 0
	return res
}

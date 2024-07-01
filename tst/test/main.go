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
		//SumOfIntervals(
		//	[][2]int{
		//		{1, 4},
		//		{3, 5},
		//		{7, 10},
		//	},
		//),
		//SumOfIntervals(
		//	[][2]int{
		//		{1, 4},
		//		{7, 10},
		//		{3, 5},
		//	},
		//),
		//
		//SumOfIntervals(
		//	[][2]int{
		//		{1, 2},
		//		{2, 3},
		//		{3, 4},
		//	},
		//),
		//SumOfIntervals(
		//	[][2]int{
		//		{0, 20}, {-100_000_000, 10}, {30, 40},
		//	},
		//),
		//SumOfIntervals(
		//	[][2]int{
		//		{300, 399},
		//		{0, 99},
		//		{100, 200},
		//		{200, 300},
		//	},
		//),

		//-92 - 73
		//SumOfIntervals(
		//	[][2]int{
		//		{22, 29}, //1
		//		{74, 90},
		//		{-92, 69}, //1
		//		{10, 37},  //1
		//		{48, 72},  //1
		//		{17, 73},  //1
		//	},
		//),

		SumOfIntervals(
			[][2]int{
				{-18, 41},
				{37, 45},
				{-21, 62},
				{86, 99},
				{-5, 7},
				{88, 95},
				{-74, -42},
				{83, 99},
				{-78, 53},
				{79, 87},
				{72, 94},
				{-80, 60},
				{74, 83},
				{70, 79},
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
					if currIntervalIndex > 0 {
						r = append(
							r,
							//intervals[:currIntervalIndex-1]...,
							intervals[:currIntervalIndex]...,
						)
					}
				} else {
					if i > 0 {
						r = append(
							r,
							intervals[:i]...,
						)
					}
				}

				//new
				newRightBorder := intervals[currIntervalIndex][1]
				if intervals[i][1] > newRightBorder {
					newRightBorder = intervals[i][1]
				}

				newLeftBorder := intervals[currIntervalIndex][0]
				if intervals[i][0] < newLeftBorder {
					newLeftBorder = intervals[i][0]
				}

				r = append(
					r,
					[2]int{
						newLeftBorder,
						newRightBorder,
					},
				)

				//middle
				if i > currIntervalIndex+1 {
					if currIntervalIndex+1 < i {
						r = append(
							r,
							intervals[currIntervalIndex+1:i]...,
						)
					}
				} else {
					if i+1 < currIntervalIndex {
						r = append(
							r,
							intervals[i+1:currIntervalIndex]...,
						)
					}
				}

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

	for _, v := range intervals {
		res += v[1] - v[0]
	}

	return res
}

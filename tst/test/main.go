package main

import "fmt"

func main() {
	fmt.Println(Partitions(5))
}

func Partitions(n int) int {

	elem := make([]int, n)

	curr := 0
	res := 0
	for i := 0; i < n; i++ {
		for f := 0; f < n; f++ {
			for j := 0; j < n; j++ {
				if curr >= j {
					elem[j]++
				}

				sum := 0
				for _, v := range elem {
					sum += v
				}

				if sum == n {
					res++
					fmt.Println(elem)
					break
				}
			}
		}

		curr++
		elem = make([]int, n)
	}

	return res
}

package main

import "fmt"

func main() {
	//sl := []int(nil)
	//nsl := append([]int{}, sl...)

	//var tst []int
	tst := make([]int, 0, 10)

	nsl := append(
		tst,
		1,
	)

	nsl = append(nsl, 2)
	//nsl = append(nsl, 3)
	nsl = append(nsl, 4)
	//nsl = append(nsl, 5)

	fmt.Println(nsl)
	fmt.Println(len(nsl))
	fmt.Println(cap(nsl))
}

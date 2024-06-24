package main

import (
	"app/tst/lib"
	"fmt"
	"time"
)

func main() {
	base := lib.CreateRandList(10, 100)

	t := time.Now()
	bTree := lib.CreateBTree(base)
	res := bTree.Sort(lib.ASC)
	fmt.Println("")
	fmt.Println(res)
	fmt.Printf("Time: %s", time.Since(t))

	//fmt.Println(values)
	//fmt.Println(res.sort())
	//fmt.Println(res.bypass())
	//fmt.Println(res.searchLower(6))
	//fmt.Println(res.searchGreater(6))
}

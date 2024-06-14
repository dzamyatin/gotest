package main

import (
	"app/tst/lib"
	"fmt"
)

func main() {
	var list lib.OneWayList
	for i := 0; i < 10; i++ {
		list.AddVal(i)
	}

	list.Flip()
	list.Flip()
	list.Flip()

	for _, v := range list.ToSlice() {
		fmt.Println(v.Value())
	}
}

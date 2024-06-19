package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

func main() {
	a := 1
	b := 2

	pa := unsafe.Pointer(&a)
	pb := unsafe.Pointer(&b)

	atomic.SwapPointer(&pa, pb)

	v := *(*int)(pa)

	fmt.Printf("%v", v)

}

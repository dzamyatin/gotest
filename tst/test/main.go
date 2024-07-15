package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "Test 1 that should be replaced 1"

	r, e := regexp.Compile(`[0-9]`)

	if e != nil {
		panic(e)
	}

	res := r.ReplaceAll([]byte(str), []byte("<>"))

	fmt.Println(string(res))
}

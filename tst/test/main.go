package main

import (
	"fmt"
	_ "golang.org/x/sync/errgroup"
)

func main() {
	fmt.Println(
		Solution("abcde"),
	)
}

func Solution(str string) []string {
	res := []string{}
	buf := []rune{}

	for _, v := range str {
		buf = append(buf, v)
		if len(buf) == 2 {
			res = append(res, string(buf))
			buf = []rune{}
		}
	}

	if len(buf) != 0 {
		res = append(res, string(buf)+"_")
	}

	return res
}

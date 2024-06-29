package main

import (
	"fmt"
	_ "golang.org/x/sync/errgroup"
)

func main() {
	fmt.Println(
		Solution("хуй"),
	)
}

func Solution(str string) []string {
	res := []string{}
	s := []rune(str)

	if len(s)%2 != 0 {
		s = append(s, '_')
	}

	for i := 1; i < len(s); i += 2 {
		res = append(res, string(s[i-1])+string(s[i]))
	}

	return res
}

func Solution4(str string) []string {
	res := []string{}

	buf := 0
	for k, v := range str {
		if k%2 == 0 {
			res = append(res, string(v))
			buf = len(res) - 1
			continue
		}
		res[buf] += string(v)
	}

	if len(res[buf]) == 1 {
		res[buf] += "_"
	}

	return res
}

func Solution3(str string) []string {
	res := [][]rune{}
	var buf []rune

	cal := 0
	for _, v := range str {
		if cal == 0 {
			buf = []rune{v, '_'}
			res = append(res, buf)
		} else {
			buf[cal] = v
			cal = -1
		}
		cal++
	}

	f := []string{}
	for _, v := range res {
		f = append(f, string(v))
	}

	return f
}

func Solution2(str string) []string {
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

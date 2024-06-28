package main

import (
	"fmt"
	_ "golang.org/x/sync/errgroup"
	"strings"
)

func main() {
	fmt.Println(
		duplicate_count("abcde"),
		duplicate_count("abcdea"),
		duplicate_count("abcdeaB11"),
		duplicate_count("indivisibility"),
	)
}

func duplicate_count(s1 string) int {
	s1 = strings.ToLower(s1)
	r := []rune(s1)
	m := make(map[rune]int, len(r))

	res := 0
	for _, v := range r {
		if c, ok := m[v]; ok && c == 1 {
			res++
		}
		m[v]++
	}

	return res
}

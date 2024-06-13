package main

import (
	"fmt"
	"sort"
)

type Test struct {
	Name string
}

type Tests []Test

func (s Tests) Len() int           { return len(s) }
func (s Tests) Less(i, j int) bool { return s[i].Name < s[j].Name }
func (s Tests) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	//m2 := [5]Test{}
	m := [4]Test{
		{"Daniil"},
		{"Alex"},
		{"Oleg"},
		{"Perez"},
	}

	test := m[:]

	//sort.Sort(test)
	sort.Slice(test, func(i, j int) bool {
		return test[i].Name < test[j].Name
	})

	//tst := m2[:]
	//copy(tst, m[:])

	fmt.Println(test)
	//fmt.Println(tst)
}

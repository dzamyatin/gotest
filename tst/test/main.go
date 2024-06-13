package main

import (
	"fmt"
	"math/rand"
)

func main() {
	test := map[string]int{
		"Hella":  0,
		"Orr":    0,
		"aweawe": 0,
	}

	for i := 0; i < 100; i++ {
		var j int32 = 0
		tgt := rand.Int31n(int32(len(test)))
		for k := range test {
			if tgt == j {
				test[k] += 1
				break

			}
			j++
		}
	}

	fmt.Println(test)
	fmt.Println(len(test))
}

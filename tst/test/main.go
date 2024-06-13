package main

import "fmt"

func main() {
	fmt.Println(test(2))
	fmt.Println(test2(2).(int))
}

type Num interface {
	int | float32 | float64 | int32 | int64 | int8
}

func test[T Num](n T) T {
	var x interface{}
	x = n

	switch x.(type) {
	case int:
		return 1
	}

	return n
}

func test2(n interface{}) interface{} {
	var x interface{}
	x = n

	switch x.(type) {
	case int:
		return 1
	}

	return n
}

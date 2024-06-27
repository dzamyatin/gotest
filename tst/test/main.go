package main

import (
	_ "golang.org/x/sync/errgroup"
)

type Test[T any] struct {
	v T
}

func (t Test[T]) get() T {
	return t.v
}

func main() {
	t := Test[string]{}
	println(
		t.get(),
	)
}

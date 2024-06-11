package main

import (
	"errors"
	"fmt"
)

type CustomErr struct {
	err error
}

func (c CustomErr) Error() string {
	//TODO implement me
	panic("implement me")
}

func (c CustomErr) Unwrap() error {
	if c.err == nil {
		return nil
	}
	return c.err
}

type MyErr struct{}

func (m MyErr) Error() string {
	//TODO implement me
	panic("implement me")
}

func main() {

	//var err error

	//fmt.Println(retErr() == nil)
	//fmt.Println(errors.Is(retErr(), MyErr{}))

	fmt.Println(errors.Is(CustomErr{err: MyErr{}}, MyErr{}))

	var e error
	fmt.Println(errors.As(CustomErr{err: MyErr{}}, &e))

	//ch := make(chan int, 1)
	//<-ch //deadlock
	//ch <- 1
	//deferFunc()
	//fmt.Println(panicFunc())
}

func retErr() error {
	return MyErr{}
}

func deferFunc() {
	i := 1
	defer func() {
		fmt.Println(i)
	}()

	i++
}

func panicFunc() (res int) {
	defer func() {
		if recover() != nil {
			fmt.Println("Done")
		}
		res = 1
	}()
	panic("Fail")
}

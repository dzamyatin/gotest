package main

import (
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

	//fmt.Println(retCustomErr() == nil)
	//fmt.Println(reflect.TypeOf(retCustomErr()))
	//fmt.Println(reflect.TypeOf(nil))
	//fmt.Println(errors.Is(retErr(), MyErr{}))

	//fmt.Println(errors.Is(CustomErr{err: MyErr{}}, MyErr{}))
	//
	//var e error
	//fmt.Println(errors.As(CustomErr{err: MyErr{}}, &e))
	//
	//fmt.Println(errors.As(CustomErr{err: MyErr{}}, &CustomErr{}))

	//ch := make(chan int, 1)
	//<-ch //deadlock
	//ch <- 1
	//deferFunc()
	//fmt.Println(panicFunc())

	s := siren{
		sval: 1,
		graps: graps{
			gval: 2,
		},
	}
	fmt.Println(
		s.test(),
	)
}

type siren struct {
	graps
	sval int
}

type graps struct {
	gval int
}

func (g graps) test() int {
	return g.gval
}

//func (s siren) test2() int {
//	return s.graps.test()
//}
//
//func (s siren) test() int {
//	return s.sval
//}

func retErr() error {
	return MyErr{}
}

func retCustomErr() error {
	var err *MyErr
	return err
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

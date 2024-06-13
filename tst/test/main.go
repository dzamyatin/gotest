package main

import (
	"app/tst/lib"
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

	//s := siren{
	//	sval: 1,
	//	graps: graps{
	//		gval: 2,
	//	},
	//}
	//fmt.Println(
	//	s.test(),
	//)

	//fmt.Println(
	//	strconv.Itoa(97),
	//)
	//
	//fmt.Println(
	//	strconv.Atoi("88"),
	//)
	//
	//fmt.Println([]rune("a")[0] == 97)
	//
	//fmt.Println('Й')
	//fmt.Println([]rune("Йух"))
	//
	//fmt.Println("----")
	//for _, v := range "Русь великая №123" {
	//	fmt.Println(v)
	//}
	//
	//fmt.Println("----")
	//for _, v := range []byte("Русь великая №123") {
	//	fmt.Println(v)
	//}

	fmt.Println("----")
	//fmt.Println(string([]byte{208, 160, 209, 131, 209, 129}))
	//fmt.Println(string([]byte{209, 131, 209, 129}))

	//conv := lib.ByteToStringConverter{}
	//conv.Add('H')
	//conv.Add('e')
	//conv.Add('l')
	//conv.Add('l')
	//conv.Add('o')
	//conv.Add(':')
	//conv.Add(208)
	//conv.Add(160)
	//conv.Add(209)
	//conv.Add(131)
	//conv.Add(209)
	//conv.Add(129)
	//conv.Add('_')
	//conv.Add('1')
	//conv.Add('2')
	//conv.Add('3')
	//conv.Add('4')
	//conv.Add('5')
	//conv.Add('6')
	//conv.Add('7')
	//conv.Add('8')
	//conv.Add('9')

	//fmt.Println("+++++++++++")
	//fmt.Println(strconv.FormatInt(209, 2))
	//fmt.Println(strconv.FormatInt(208, 2))
	//fmt.Println(strconv.ParseInt("11010001", 2, 64))
	//fmt.Println(strconv.ParseInt("11010000", 2, 64))
	//fmt.Println(0b11010001 & 0b11010000)
	////return
	//fmt.Println("-->")
	//fmt.Println(string([]byte{208, 159}))
	//fmt.Println(string([]byte{209, 128}))
	//fmt.Println(string([]byte{208, 184}))
	//fmt.Println(string([]byte{208, 178}))
	//fmt.Println("<--")

	conv := lib.ByteToStringConverter{}

	bt := " Привет! Как дела? 😀 Norn, sam kak&? 123"
	for _, v := range []byte(bt) {
		//fmt.Println(v)
		//fmt.Println(string([]byte{v}))
		conv.Add(v)
	}

	fmt.Println(string(conv.Convert()))
	fmt.Println(string(conv.ConvertAll()))
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

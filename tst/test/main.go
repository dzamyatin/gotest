package main

import "log"

func main() {
	tst := 1
	f := func(tst int) {
		log.Println(tst)
	}
	defer f(tst)
	tst = 2

	//arr := [3]int{1, 2, 3}
	//slice := []int{1, 2, 3}
	//m := map[int]int{1: 1, 2: 2, 3: 3}
	//
	//e(arr, slice, m)
	//
	//t := [3]int(slice)
	//
	//tst(t[:]...)
	//
	//fmt.Println([3]int(slice) == arr) //true
	//fmt.Println(slice == slice) //build failed
	//fmt.Println(m == m) //build failed

	//p := Person{Name: "Alice", Age: 30}
	//p2 := Person{Name: "Alice", Age: 30}
	//e := Employee{Name: "Alice", Age: 30, ID: 123}

	//log.Println(p == p2)

	//var i interface{} = &p
	//change2(&i)

	//change4(&p)

	//log.Println(p.Age)
}

//type IPers interface {
//	setAge(age int)
//}
//
//type Person struct {
//	Name string
//	Age  int
//	//Ptr  []int
//	//Ptr map[int]int
//}
//
//func (p *Person) setAge(age int) {
//	p.Age = age
//}
//
//type Employee struct {
//	Name string
//	Age  int
//	ID   int
//}
//
//func tst(a ...int) {
//	fmt.Println(a)
//}
//func change(p Person) {
//	p.Age = 31
//}
//
//func change2(p *interface{}) {
//	t := (*p).(*Person)
//
//	t.setAge(31)
//}
//
//func change3(p IPers) {
//	p.setAge(31)
//}
//
//func change4(p IPers) {
//	g := p.(*Person)
//	g.setAge(31)
//}

func e(...interface{}) {

}

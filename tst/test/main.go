package main

import (
	"log"
)

type Person struct {
	Name string
	Age  int
	//Ptr  []int
	//Ptr map[int]int
}

func (p *Person) setAge(age int) {
	p.Age = age
}

type Employee struct {
	Name string
	Age  int
	ID   int
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	//p2 := Person{Name: "Alice", Age: 30}
	//e := Employee{Name: "Alice", Age: 30, ID: 123}

	//log.Println(p == p2)

	change2(p)

	log.Println(p.Age)
}

func change(p Person) {
	p.Age = 31
}

func change2(p interface{}) {
	t := p.(Person)
	t.Age = 31
	t.setAge(31)
}

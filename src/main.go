package main

import (
	"fmt"
)

type Mutatable struct {
	a int
	b int
	name test
}

type test struct {
	name string
}

func (m Mutatable) StayTheSame(n1 int, n2 int) {
	m.a = n1
	m.b = n2
}

func (m *Mutatable) Mutate(n1 int, n2 int) {
	m.a = n1
	m.b = n2
}
func (m *Mutatable) Mutate1(test test) {
	m.name = test
}

func main() {
	m := &Mutatable{0, 0, test{name: "aaa"}}
	fmt.Println(m)
	m.StayTheSame(5,7)
	fmt.Println(m)
	m.Mutate(6,8)
	fmt.Println(m)
	m.Mutate1(test{name: "bob"})
	fmt.Println(m)
}
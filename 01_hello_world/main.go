package main

import "fmt"

type Myint int

func (m Myint) Abs() int {
	if m < 0 {
		return int(-m)
	}
	return int(m)
}

func main() {
	a := 33
	fmt.Printf("%T\n", a)

	var m Myint = -10

	fmt.Printf("%T, %v\n", m, m)
	fmt.Printf("%v\n", m.Abs())
}

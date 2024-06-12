package main

import "time"

type Abser interface {
	Div() float64
}

type Vertex struct {
	X, Y int
}

type Person struct {
	Name string
	Age  int
}

type IPAddr [4]byte

type MyError struct {
	When time.Time
	What string
}

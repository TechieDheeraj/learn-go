package main

import (
	"fmt"

	l "github.com/TechieDheeraj/learn-go/basic_ds/list"
	"golang.org/x/tour/reader"
)

func main() {
	fmt.Println("Hello World")
	basicStuff()
	reader.Validate(MyReader{})
	checkFuncs()
	var linkedList l.List

	linkedList.PushFront(10)
	linkedList.PushFront(11)
	linkedList.PushFront(12)
	linkedList.Print()
}

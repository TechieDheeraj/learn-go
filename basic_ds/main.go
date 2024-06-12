package main

import (
	"fmt"

	t "github.com/TechieDheeraj/learn-go/basic_ds/bst"
	gr "github.com/TechieDheeraj/learn-go/basic_ds/graph"
	l "github.com/TechieDheeraj/learn-go/basic_ds/list"
	"golang.org/x/tour/reader"
)

func main() {
	fmt.Println("Hello World")
	basicStuff()
	reader.Validate(MyReader{})
	checkFuncs()
	ll := l.New()

	ll.PushFront(10)
	ll.PushFront(11)
	ll.PushFront(12)
	ll.Print()

	bst := t.Init(100)
	bst.Insert(50)
	bst.Insert(110)
	bst.Insert(130)
	bst.Insert(30)
	bst.Insert(40)
	bst.Insert(109)
	bst.Print()

	g := &gr.Graph{}
	n1 := gr.NewVertex(1)
	n2 := gr.NewVertex(2)
	n3 := gr.NewVertex(3)
	n4 := gr.NewVertex(4)
	n5 := gr.NewVertex(5)
	n6 := gr.NewVertex(6)

	g.AddVertex(n1)
	g.AddVertex(n2)
	g.AddVertex(n3)
	g.AddVertex(n4)
	g.AddVertex(n5)
	g.AddVertex(n6)

	g.AddEdge(n5, n3)
	g.AddEdge(n3, n1)
	g.AddEdge(n1, n2)
	g.AddEdge(n3, n4)
	g.AddEdge(n2, n4)
	g.AddEdge(n2, n6)

	g.Print()

}

package list

import "fmt"

type Node struct {
	next *Node
	val  any
}

type List struct {
	root Node
	len  int
}

func (l *List) Init() *List {
	l.root.next = nil
	l.root.val = 0
	l.len = 0
	return l
}

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func New() *List {
	return new(List).Init()
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *Node {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) insert(e, at *Node) *Node { // insert element (e) after (at)
	// at -> e
	e.next = at.next
	at.next = e
	l.len++
	return e
}

// insertValue is wrapper of insert(&Node{val: v}, at)
func (l *List) insertValue(v any, at *Node) *Node {
	return l.insert(&Node{val: v}, at)
}

func (l *List) PushFront(v any) *Node {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

func (l *List) Print() {
	var printNode func(n *Node)

	printNode = func(n *Node) {
		if n == nil {
			return
		}
		fmt.Printf("%d->", n.val)
		printNode(n.next)
	}

	printNode(l.root.next)
	fmt.Println()
}

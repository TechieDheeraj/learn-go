package bst

import (
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func Init(val int) *TreeNode {
	var node TreeNode
	node.val = val
	node.left = nil
	node.right = nil
	return &node
}

func (t *TreeNode) Insert(k int) {
	if t.val < k {
		if t.right == nil {
			t.right = Init(k)
		} else {
			t.right.Insert(k)
		}
	} else {
		if t.left == nil {
			t.left = Init(k)
		} else {
			t.left.Insert(k)
		}
	}
}

func (t *TreeNode) Search(k int) bool {
	if t == nil {
		return false
	}

	if t.val < k {
		t.right.Search(k)
	} else {
		t.left.Search(k)
	}
	return true
}

func (t *TreeNode) Print() {
	if t == nil {
		return
	}

	t.left.Print()
	fmt.Printf("%d->", t.val)
	t.right.Print()
}

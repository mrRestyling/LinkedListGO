package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type BSTree struct {
	Root *Node
}

func main() {
	bst := &BSTree{}
	bst.Insert(5)
	bst.Insert(666)
	bst.Insert(667)
	bst.Insert(6)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(6)
	bst.Round(bst.Root)
}

func (bst *BSTree) Insert(data int) {
	if bst.Root == nil {
		bst.Root = &Node{Val: data}
	} else {
		bst.insert(data, bst.Root)
	}
}

func (bst *BSTree) insert(data int, current *Node) {
	if data < current.Val {
		if current.Left == nil {
			current.Left = &Node{Val: data}
		} else {
			bst.insert(data, current.Left)
		}
	} else {
		if current.Right == nil {
			current.Right = &Node{Val: data}
		} else {
			bst.insert(data, current.Right)
		}
	}
}

func (bst *BSTree) Round(root *Node) {
	if root != nil {
		bst.Round(root.Left)
		fmt.Println(root.Val)
		bst.Round(root.Right)
	}
}

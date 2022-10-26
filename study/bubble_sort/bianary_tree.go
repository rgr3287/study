package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Left  *Node
	Data  int
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func main() {
	var n1, n2, n3, n4, n5, n6, n7 Node
	n1 = Node{Left: &n2, Data: 50, Right: &n3}
	n2 = Node{Left: &n4, Data: 25, Right: &n5}
	n3 = Node{Left: &n6, Data: 75, Right: &n7}
	n4 = Node{Data: 10}
	n5 = Node{Data: 33}
	n6 = Node{Data: 56}
	n7 = Node{Data: 89}

	tree := BinaryTree{Root: &n1}
	tree.Insert(90)
	tree.Insert(1)
	tree.Insert(5)
	tree.Insert(74)
	tree.Insert(65)
	fmt.Println(tree.Search(66))
	//fmt.Println(tree.Root.Right.Right.Right)
	//fmt.Println(tree.Root.Left.Left)
	resultNode(tree.Root)
}

func resultNode(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.Data)
	resultNode(n.Left)
	resultNode(n.Right)
}

func (b *BinaryTree) Search(target int) bool {
	if b.Root.Data == target {
		return true
	}
	if b.Root.Left != nil {
		leftBinary := BinaryTree{Root: b.Root.Left}
		if leftBinary.Search(target) {
			return true
		}
	}
	if b.Root.Right != nil {
		rightBinary := BinaryTree{Root: b.Root.Right}
		if rightBinary.Search(target) {
			return true
		}
	}
	return false
}

func (b *BinaryTree) Insert(newVal int) error {
	if newVal <= b.Root.Data {
		if b.Root.Left == nil {
			b.Root.Left = &Node{Data: newVal}
		} else {
			leftBinary := BinaryTree{Root: b.Root.Left}
			leftBinary.Insert(newVal)
		}
	} else {
		if b.Root.Right == nil {
			b.Root.Right = &Node{Data: newVal}
		} else {
			RightBinary := BinaryTree{Root: b.Root.Right}
			RightBinary.Insert(newVal)
		}
	}
	return errors.New("insert overflow")
}

func (b *BinaryTree) Delete(target int) error {
	return nil
}

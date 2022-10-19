package main

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	First  *Node
	Last   *Node
	Length int
}

type Node struct {
	Prev *Node
	Data int
	Next *Node
}

func main() {
	var node1, node2, node3 Node

	node1 = Node{nil, 1, &node2}
	node2 = Node{&node1, 2, &node3}
	node3 = Node{&node2, 3, nil}
	linkedList := LinkedList{&node1, &node3, 3}
	fmt.Println(linkedList.Get(0))
	fmt.Println(linkedList.Get(1))
	fmt.Println(linkedList.Get(2))

	fmt.Println(linkedList.First.Data, linkedList.Delete())
	fmt.Println(linkedList.First.Data, linkedList.Delete())
	fmt.Println(linkedList.First.Data, linkedList.Delete())
}

func (list *LinkedList) Get(index int) (int, error) {
	now := list.First
	if index < 0 || index >= list.Length {
		return 0, errors.New("index not valid")
	}
	for i := 0; i < index; i++ {
		next, err := Next(now)
		if err != nil {
			return 0, err
		}
		now = next
	}

	return now.Data, nil
}

func Next(node *Node) (*Node, error) {
	if node.Next == nil {
		return nil, errors.New("last is nil")
	}

	return node.Next, nil
}

func (list *LinkedList) Delete() bool {
	if list.First.Next == nil {
		return false
	}
	list.First = list.First.Next
	return true
}

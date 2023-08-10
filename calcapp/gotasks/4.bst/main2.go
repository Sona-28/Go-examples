package main

import (
	"fmt"
)

type Node1 struct {
	key   int
	left  *Node1
	right *Node1
}

func (n *Node1) insert(root *Node1, i int, r chan *Node1) {
	// fmt.Println("insert", i)
	if root == nil {
		r <- &Node1{i, nil, nil}
		fmt.Println("insert",i)

	} else {
		if i < root.key {
			go n.insert(root.left, i, r)
			// root.left = <-r
		} else {
			go n.insert(root.right, i, r)
			// root.right = <- r
		}
	}
}

func main() {
	root := &Node1{}
	// fmt.Println(root)

	result := make(chan *Node1)
	arr := [5] int {20,8,50,7,9}
	for i := 0; i < len(arr); i++ {
		go root.insert(root, arr[i], result)
		// result <- root
		root = <-result

		fmt.Println(root)
	}
	// fmt.Print(result)
	// fmt.Println(root)

	close(result)
}

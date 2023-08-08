package main

import (
	"fmt"
	"time"
)

type node struct{
	key int
	left *node
	right *node
}



func (n *node) insert(root *node, k int) *node{
	if root == nil{
		newnode := node{k, nil, nil}
		fmt.Println("The element is inserted ",k)
		return &newnode
	}else{
		if k < root.key{
			root.left = n.insert(root.left,k)
		}else{
			root.right = n.insert(root.right,k)
		}
	}
	return root
}

func (n *node) search(cur *node, k int, parent *node) *node{
	start := time.Now()
	for ;cur!=nil && cur.key!=k;{
		parent = cur
		if k<cur.key{
			cur = cur.left
		}else{
			cur = cur.right
		}
	}
	timeElapsed := time.Since(start)
	fmt.Printf("The time taken for the searching is %s\n", timeElapsed)
	return cur
}

func main(){
	var root *node = nil
	root = root.insert(root,49)
	root = root.insert(root,46)
	root = root.insert(root,80)
	root = root.insert(root,48)
	root = root.insert(root,30)
	root = root.insert(root,50)
	r := root.search(root, 46, nil)
	fmt.Print(r)
}
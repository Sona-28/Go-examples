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
	}
		if k < root.key{
			root.left = n.insert(root.left,k)
		}else{
			root.right = n.insert(root.right,k)
		}
	
	return root
}

func (n *node) search(cur *node, k int) *node{
	start := time.Now()
	for ;cur!=nil && cur.key!=k;{
		// parent = cur
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
	root:= &node{}
	//var root *node = nil
	root = root.insert(root,10)
	root = root.insert(root,30)
	root = root.insert(root,5)
	root = root.insert(root,31)
	root = root.insert(root,40)
	// root = root.insert(root,50)
	r := root.search(root, 10)
	fmt.Print(r.key,r.left.key, r.right.key)
}
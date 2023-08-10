package main

import (
	"fmt"
	
)

type Node struct{
	key int
	left *Node
	right *Node
}



func (n *Node) insert(r chan Node, k int) {
	fmt.Println("insert")
	if root == nil{
		newNode := &Node{k, nil, nil}
		fmt.Println("The element is inserted ",k)
		result <- newNode
	}else{
		if k < root.key{
			go n.insert(root.left,k, result)
			root.left = <- result
		}else{
			go n.insert(root.right,k, result)
			root.right = <- result

		}
	}
	r <- Node{1,nil,nil}
	fmt.Println(r)
}


func main(){
	// var wg sync.WaitGroup
	root:= Node{}
	result := make(chan Node)
	arr := [5]int {10,30,45,20,5}
	for i:=range(arr){
		go root.insert(result,arr[i])
		root = <- result
	}
	// r := root.search(root, 10)
	// fmt.Print(r)
	close(result)
}
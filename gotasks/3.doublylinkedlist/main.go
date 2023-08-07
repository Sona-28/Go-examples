package main

import "fmt"

type Node struct{
	data int
	prev *Node
	next *Node
	

}

var head_node *Node
var tail_node *Node

func (n *Node) insert(nd int){
	if
}



// func main(){
// 	head_node := Node{}
// 	second_node := Node{4, &head_node, nil}
// 	head_node.next = &second_node

// 	for i:=1; i!=0; {
// 		fmt.Printf("%d ", head_node.data)
// 		head_node = *head_node.next
// 	}
	// fmt.Println(head_node)
// }
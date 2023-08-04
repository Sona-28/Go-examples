package main

import "fmt"

func main(){
	// var l map[string] int
	m:= make(map[string] int)
	m["salary"] = 2000
	m["age"] = 20
	m["EId"] = 1

	if m == nil{
		fmt.Println("m is nil")
	} else{
		fmt.Println(m)
	}

	// if l == nil{
	// 	fmt.Println("l is nil")
	// } else{
	// 	fmt.Println(l)
	// }
}
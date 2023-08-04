package main

import "fmt"

func main(){
	var a int = 10
	var p *int = &a
	fmt.Println(a)
	fmt.Println(p)

	*p = 2000
	fmt.Println(a)
	fmt.Println(p)

}
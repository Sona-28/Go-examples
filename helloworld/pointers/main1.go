package main

import "fmt"

func main(){
	var a int = 10
	var p  = &a
	var pp = &p
	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(pp)


	*p = 2000
	fmt.Println(a)
	fmt.Println(p)

}
package main

import "fmt"

func main(){
	defer fmt.Println("Last0")

	fmt.Println("First")
	defer fmt.Println("Last")
	defer fmt.Println("Last1")
	defer fmt.Println("Last2")

	fmt.Println("Second")
}
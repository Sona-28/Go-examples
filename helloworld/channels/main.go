package main

import "fmt"

func routine(ch chan int){
	fmt.Println(100+ <-ch)
}

func main(){
	ch_demo :=make(chan int)
	go routine(ch_demo)
	ch_demo <- 234

	fmt.Println("Value of channel ",ch_demo)
	fmt.Printf("Type of channel %T",ch_demo)
}
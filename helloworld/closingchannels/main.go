package main

import "fmt"

func myfunc(channel chan string){
	for i:=0;i<10;i++{
		channel <- "Go lang is awesome"
	}
	close(channel)
}

func main(){
	c:=make(chan string, 8)
	go myfunc(c)
	counter := 0
	for{
		res,ok := <- c
		counter++
		if !ok{
			fmt.Println("close",ok)
			break
		}
		fmt.Println("Open",counter, res, ok)
	}
}
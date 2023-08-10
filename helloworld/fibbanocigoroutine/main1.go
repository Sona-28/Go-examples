package main

import (
	"fmt"
	"time"
)

func fibbanoci(n int, c chan int){
	x,y := 0,1
	for i:=0;i<n;i++{
		c <- x
		x,y =y,x+y
	}
	close(c)
}

func main(){
	start := time.Now()
	n := make(chan int,50)
	go fibbanoci(cap(n), n)
	// for i:=range n{
	// 	fmt.Println(i)
	// }
	end := time.Now().Sub(start)
	fmt.Println(end)
}
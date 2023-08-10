package main

import (
	"fmt"
	"time"
)

func fibbanoci(i int) int{
	if i<=1{
		return i
	}else{
		return fibbanoci(i-1)+fibbanoci(i-2)
	}
}

func main(){
	start := time.Now()
	n := 50
	for i:=0;i<n;i++{
		fibbanoci(i)
		// fmt.Println(fibbanoci(i))
	}
	end := time.Now().Sub(start)
	fmt.Println(end)
}
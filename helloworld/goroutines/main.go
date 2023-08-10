package main

import (
	"fmt"
	"time"
)

func hello(){
	fmt.Println("In go routine")
}

func main(){
	fmt.Println("In main")
	fmt.Println("Calling hello")
	go hello()
	time.Sleep(10*time.Microsecond)
	fmt.Println("Main completed")
}
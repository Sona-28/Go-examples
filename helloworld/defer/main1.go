package main

import "fmt"

func one(){
	defer two()
	fmt.Println(1)

}

func two(){
	defer three()
	fmt.Println(2)

}

func three(){
	defer fmt.Println(3)

}

func main(){
	defer one()
}
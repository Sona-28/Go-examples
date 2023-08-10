package main

import "fmt"

func main(){
	var num int
	fmt.Print("Enter the number: ")
	fmt.Scanln(&num)
	for i:=1; i<num+1; i++{
		for j:=0;j<num-i+1;j++{
			fmt.Print(" ")
		}
		c:=1
		for j:= 1; j<i+1; j++{
			fmt.Print(" ",c)
			c = c*(i-j)/j
		}
		fmt.Println()
	}
}
package main

import "fmt"

func main(){
	pm := []int{2,3,5,7,11,13,17,19,23,29}
	for index, num := range pm{
		fmt.Printf("Primenumber(%d) = %d\n",index+1, num)
	}
}
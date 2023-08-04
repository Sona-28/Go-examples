package main

import (
	"fmt"
	"math"
)

func calc(a, b float64)(float64,float64){
	sum :=a+b
	diff := math.Abs(a-b)
	return sum,diff
}

func main(){
	var a float64=0
	var b float64=20
	sum, diff:=calc(a,b)
	fmt.Printf("Sum: %.1f, Difference: %.1f",sum,diff)
}
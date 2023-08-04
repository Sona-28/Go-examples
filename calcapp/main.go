package main

import (
	"fmt"
	"netxd/calcapp/calc"
)

func main(){
	a:=10
	b:=20
	sum :=calc.Add(a,b)
	fmt.Println(sum)
}


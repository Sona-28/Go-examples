package main

import "fmt"

func main(){
	var alp = map [string] int{
		"a":1,
		"b":2,
		"c":3,
		"d":4,
	}
	alp["e"] = 5
	alp["f"] = 6

	x,y := alp["g"];

	fmt.Println(x)
	fmt.Println(y)

	delete(alp, "f")
	fmt.Println(alp)
}
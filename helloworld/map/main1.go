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

	for alphabets, num := range alp{
		fmt.Println(alphabets, num)
	}
}
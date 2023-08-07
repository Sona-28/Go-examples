package main

import "fmt"

func main(){
	str := "ABCDEFGHIJKL"
	runes := []rune(str)

	for i := 2; i<len(runes);i+=2{
		runes[i], runes[i+1] = runes[i+1], runes[i]
		i+=2
	}
	fmt.Println(string(runes))
}

package main

import "fmt"

type myString string

func (mystr myString) reverse() string{
	s:= string(mystr)
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i<j;i,j=i+1,j-1{
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main(){
	mystr := myString("OLLEH")
	fmt.Println(mystr.reverse())

}
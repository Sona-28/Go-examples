package main

import "fmt"

func main(){
	var s = [5]string{"Harry","Louis","Liam","Zayn","Niall"}
	var a []string = s[1:4]
	fmt.Println("Array s = ", s)
	fmt.Println("Slice a = ", a)

	slice1:= s[1:4]
	slice2 := s[:3]
	slice3 := s[2:]
	slice4:=s[:2]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	
	
}
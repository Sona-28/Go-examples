package main

import "fmt"

func main(){
	slice1 := []string{"C","C++","Java"}
	slice2 := append(slice1, "Python", "Ruby", "Go")

	fmt.Println(slice1, len(slice1),cap(slice1))
	fmt.Println(slice2,len(slice2),cap(slice2))

	slice2[1] = "Rust"
	fmt.Println(slice1, cap(slice1))
	fmt.Println(slice2)
}
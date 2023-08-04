package main

import "fmt"

type Student struct{
	RollNumber int
	Name string
}

func main(){
	s:= Student{11, "Jack"}

	ps:= &s
	fmt.Println(ps)

	fmt.Println((*ps).Name)
	fmt.Println(ps.Name)

	ps.RollNumber = 31
	fmt.Println(ps)

}
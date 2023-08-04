package main

import "fmt"

func main(){
	a := [7]string{"Mon","Tue","Wed","Thurs","Fri","Sat","Sun"}
	slice1 :=a[1:]
	slice2 := a[3:]

	fmt.Println("----------Before Modifications------------")
	fmt.Println("a = ", a)
	fmt.Println("slice1 = ",slice1)
	fmt.Println("slice2 = ", slice2)

	slice1[0] = "TUE"
	slice1[1] = "WED"
	slice1[2] = "THURS"

	slice2[1] = "FRIDAY"

	fmt.Println("----------After Modifications------------")
	fmt.Println("a = ", a)
	fmt.Println("slice1 = ",slice1)
	fmt.Println("slice2 = ", slice2)
}
package main

import "fmt"

func main(){
	a := []int{50,25,2,75,45}
	n := len(a)
	for i:=0; i<n-1; i++{
		for j:=0; j<n-i-1; j++{
			if a[j]>a[j+1]{
				a[j],a[j+1] = a[j+1],a[j]
			}
			
		}
	}
	fmt.Println(a)

}
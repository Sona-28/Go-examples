package main

import "fmt"

func main(){
	var a = [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	var i,j int
	for i=0;i<3;i++{
		for j=0;j<3;j++{
			fmt.Printf("%d ",a[i][j])
		}
		fmt.Println()
	}
}
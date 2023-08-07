package main

import "fmt"

func main(){
	var str string
	count :=0
	fmt.Scanln(&str)
	for i:=0;i<len(str)-1;i++{
		if string(str[i])=="g"{
			if string(str[i+1])=="o"{
				count++
			}
		}
	}
	fmt.Println(count)

}
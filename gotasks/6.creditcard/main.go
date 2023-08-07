package main

import (
	"fmt"

	c "github.com/akriventsev/go-card"
)

func main(){
	_,err := c.NewCard("4716339239466898", "334", 12, 2023, "Ivanov Ivan")

	if err!= nil {
		fmt.Print(err)
	}else{
		fmt.Println("Your card is valid")
	}
}
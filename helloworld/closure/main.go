package main

import "fmt"

func main() {
	// sum := add(10,20)
	fmt.Println(add(10, 20)(), add(40, 20)())
	// fmt.Println(add(40, 20)())
}
func add(a, b int) func() int {
	c := 100
	return func() int {
		return a + b + c
	}
}

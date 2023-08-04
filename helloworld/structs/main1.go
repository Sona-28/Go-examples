package main

import "fmt"

type Point struct{
	X,Y float64
}

func (p Point) IsAbove(y float64) bool{
	return p.Y>y
}

func main(){
	p := Point{2.0, 4.0}
	fmt.Println(p)
	fmt.Println(p.IsAbove(5))

}
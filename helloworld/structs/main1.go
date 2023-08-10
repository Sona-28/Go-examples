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
	ps := &p

	fmt.Println(ps)
	fmt.Println( p.IsAbove(3.0))

}
package main

import "fmt"

type Point1 struct{
	X,Y float64
}

func (p *Point1) Translate(dx, dy float64){
	p.X = p.X+dx
	p.Y = p.Y+dy
}

func main(){
	p := Point1{2.0, 4.0}
	fmt.Println(p)
	p.Translate(9,7)
	fmt.Println(p)

}
package main

import (
	"fmt"
	 
)

func main() {
	fmt.Println("------------")

	t:=Triangle{width: 10,height: 5}
	t2:=NewTriangle( 10, 50)

	fmt.Println(t.Area())
	fmt.Println(t2.Area())
	fmt.Println(ShapeArea(&t))
	fmt.Println(ShapeArea(t2))
	fmt.Println(t)
	fmt.Println(&t)
	 
 

}

type Triangle struct{
	width float32
	height float32

}

func NewTriangle(w float32, h float32) *Triangle{
	return &Triangle{width: w,height: h}
}

func (t *Triangle) Area() float32{
	return (t.height*t.width)/2
}


type Shape interface{
	Area() float32
 
}

func ShapeArea(s Shape) float32{
	return s.Area()
}


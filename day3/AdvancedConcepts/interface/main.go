package main

import(
	"fmt"
)



func main(){

	c := Circle{Radius:2.458}
	PrintShapeInfo(c)

}


type Shape interface{
	Area() float64
	Perimeter() float64
}

type Circle struct{
	Radius float64
}



func (s Circle) Area() float64{
	return 3.14 * s.Radius * s.Radius
}


func (s Circle) Perimiter() float64{
	return 2*3.14*s.Radius
}
func PrintShapeInfo(s Shape){
	fmt.Prinln("Area :" s.Area())
	fmt.Prinln("Perimiter :" s.Perimeter())

}
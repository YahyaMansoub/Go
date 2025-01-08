package main

import (
	"fmt"
	"math"
)


type Shape interface{
	Area() float64
	Perimeter() float64
}

type Rectangle struct{
	Width float64
	Height float64
}
func (r Rectangle) Area() float64{
	return r.Width*r.Height
}
func (r Rectangle) Perimeter() float64{
	return 2*r.Width+2*r.Height
}
type Circle struct{
	Radius float64
}
func (c Circle) Area() float64{
	return 3.14*3.14*c.Radius
}
func (c Circle) Perimeter() float64{
	return 2*3.14*c.Radius
}
type Triangle struct{
	SideA float64
	SideB float64
	SideC float64
}

func (t Triangle) Perimeter() float64{
	return t.SideA+t.SideB+t.SideC
}
func (t Triangle) Area() float64{
	a:=t.SideA
	b:=t.SideB
	c:=t.SideC
	var s float64 =( a + b + c )/2
	return math.Sqrt(s*(s-a)*(s-b)*(s-c))
}
func calculateArea(shape interface{}) {
    switch s := shape.(type) {
    case Circle:
        fmt.Printf("Circle area: %.2f\n", s.Area())
    case Rectangle:
        fmt.Printf("Rectangle area: %.2f\n", s.Area())
	case Triangle:
        fmt.Printf("Triangle area: %.2f\n", s.Area())
    default:
        fmt.Println("Unknown shape")
    }
}
func main(){
	c:= Circle{Radius:2.14}
	fmt.Println("The Area of the Circle is:",c.Area())
	fmt.Println("The Perimeter of the Circle is:",c.Perimeter())

	r:= Rectangle{Height:45.12,Width:13.12}
	fmt.Println("The Area of the Circle is:",r.Area())
	fmt.Println("The Perimeter of the Circle is:",r.Perimeter())

	t:= Triangle{SideA:15.12,SideB:2.1,SideC:6.12}
	fmt.Println("The Area of the Circle is:",t.Area())
	fmt.Println("The Perimeter of the Circle is:",t.Perimeter())

}
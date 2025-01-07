package main

import "fmt"

const pi float64 = 3.14

func main(){
	var x,y int
	x=10
	y=20
	var name="Golang"
	fmt.Printf("x = %d, y = %d, %s\n",x,y,name)

	x:=10

	const pi float64 = 3
	fmt.Println(pi)

}
package main

import (
	"log"
	"um6p.ma/functions"
	"math.com/utils"
)


func main(){

	a,b:=4,8
	log.Println("the area is",functions.Rectarea(a,b))
	log.Println("the area is",functions.Rectperim(a,b))
	log.Println("the area of square is",utils.square(a,b))



	
}
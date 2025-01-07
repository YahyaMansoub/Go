package main

import (
	"log"
	"main/math1"
	"packagemanag/maththings/utils"
)


func main(){
	a,b:=2,4
	log.Println(a,b)
	result := math1.Mult(a, b)
    log.Println(a, b, result)
	result := utils.Square(a) 
    log.Println("The square of", a, "is", result)

}
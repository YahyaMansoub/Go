package main

import (
	"fmt"
	"time"
)


func GenSlic() []int{
	var slice []int 
	var i int
	for i = range 10{
		slice = append(slice,i+1)	
	} 
	return slice
}

func squareNumber(number int){
	result:=number*number
	fmt.Println("Square of",number,"is",result,".")
}


func main(){
	slice := GenSlic()
	for _,value:= range slice{
		go squareNumber(value)
	}
	time.Sleep(time.Second)

}
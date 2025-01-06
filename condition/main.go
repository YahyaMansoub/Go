package main

import "fmt"

func main(){
	var length float64
	var width float64


	if length == 0 {
		fmt.Println("length empty")
	}
	else if width == 0 {
		fmt.Println("width empty")
	}
	else if isLengthEmpty:= length == 0; isLengthEmpty{  //the variable is valable only within the if statement , undefined out fo the if statement 
		fmt.Println("length empty")  

	}

}
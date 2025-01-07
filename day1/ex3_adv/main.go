package main

import (
	"fmt"
	"con"
)

func main(){
	var quit bool = false
	while !quit {
		var input string
		var grade []float64
		fmt.Println("Welcome!\n")
		fmt.Scan("|-->",&input)
		if input == "quit" {
			fmt.Println("Good Bye!")
			quit = true
		}else {
			var num float64
			fmt.Scan("|-->",&num)
			num, err := strconv.ParseFloat(str, 64)
	        if err != nil {
		         fmt.Println("Error converting string to float:", err)
	        } else {
		      fmt.Println("Converted number:", num)
	        }
		    }

	}
}
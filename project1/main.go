package main

import (
	"fmt"
	"math/rand"
)

func main(){


	
	value  := rand.Int63n(101) 
	
	fmt.Println("the random number: ", value)


    var input string
	var quit bool = false
	var score int
	for !quit {
		fmt.Println("Welcome to the guessing game!\n")
		
		fmt.Println("Do you want to play?(Yes(Y)/No(N)), your score is",score)
		if input == "y" || input == "Y" {
			var limit int
			fmt.Println("Enter how much tries you are allowed to have:")
			fmt.Scan(&limit)
			for limit<1{
				fmt.Println("Please enter a correct limit.(limit<1)")
			}

			   fmt.Scan(&input)
			   var cat int
			   fmt.Println("define range manually(1) or choose category(0)")
			   fmt.Scan(&cat)
			   if cat == 0{
				fmt.Println("Choose a category(easy/medium/hard):")
			    fmt.Scan(&input)
				
			   }
		       fmt.Println("Define range:")
		       var ub int64
		       fmt.Scan(&ub)
		       for ub<1{
			       fmt.Println("Invalid upper bound , it should be bigger than 1")
		       }
		
		       fmt.Println("The game starts now...\n")
		       value := rand.Int63n(ub+1)
		       var uservalue int64

		       fmt.Println(value)
		       fmt.Scan("|-->",&uservalue)
		       for uservalue != value && limit > 0{
			       if uservalue < value {
				           fmt.Println("Your Value is Lower!\n")
						   limit--
				           fmt.Scan(&uservalue)
			       }else if uservalue > value && limit > 0{
				           fmt.Println("Your Value is Higher!")
						   limit--
				           fmt.Scan(&uservalue)
			       }
		          }
				if limit>=1{
					score++
					fmt.Println("You have won this round!\n")
					fmt.Println("Your score now is ",score)
				}else{
					fmt.Println("You have lost this round!\n")
					fmt.Println("Your score now is ",score)
				}
		
		
			
		}else if input == "n" || input == "N"{
			fmt.Println("Good bye then !!")
		} else{
			fmt.Println("Invalid input. Please enter Y or N")
		}
		
		
	
		

	}




}
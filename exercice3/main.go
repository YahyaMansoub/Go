package main

import "fmt"



func checkgrade(grade float64){
	if grade >= 100 || grade < 0 {
		fmt.Println("Invalid grade")
	}
	if grade >= 90 {
			fmt.Println("Grade 1 got : A ,",grade)
		}else if grade >= 80 && grade<90 {
			fmt.Println("Grade 1 got : B ,",grade)
		}else if grade >= 70 && grade<80 {
			fmt.Println("Grade 1 got : C ,",grade)
		}else if grade >= 60 && grade<70 {
			fmt.Println("Grade 1 got : D ,",grade)
		}else if grade >= 90 {
			fmt.Println("Grade 1 got : A ,",grade)
		}else if grade >= 0 && grade<60 {
			fmt.Println("Grade 1 got : F ,",grade)
		}

}

func main(){

	// first assignement1
	var grade [5]float64
	var sum float64=0.0
	

	for i :=0; i <5; i++{

		var input float64

        
		fmt.Println("Enter Grade ",i+1)
		fmt.Scan(&input)

		if grade[i]<=100 && grade[i]>=0{

			grade[i]=input
			sum+=input

		}else{
			fmt.Println("Invalid grade, retry again :")
			i--
		}

	}
	fmt.Println("the average of the grades is:",sum/5)
	// second assignment2
	for i :=0; i<5; i++{

		checkgrade(grade[i])

	}
	// third assignement3

	var quit bool = false
	
	for !quit{
		var choice string
		var input float64

		fmt.Println("Do you want to continue(Yes(Y)/No(N))?")
		fmt.Println("-->")78

		fmt.Scan(&choice)
		if choice == "Y" || choice == "y" {
			fmt.Println("Enter Grade")
			fmt.Println("-->")
			fmt.Scan(&input)
			checkgrade(input)
		}else if choice == "N" || choice == "n" {
			quit = true
		}else{
			fmt.Println("Write an adequate prompt (Y/N) or (y/n) please.")
		}

	}



}
package main


import (
	"log"
)

func main(){
	numbers := [5]int{}
	numbers_premium := [5][5]int{}

	log.Println("the length:",len(numbers))
	log.Println("the capacity:",cap(numbers))


	for i:=0;i< len(numbers);i++{
		log.Println(i," ",numbers[i])
	}
	for i:=0;i< len(numbers_premium);i++{
		for j:=0;j<len(numbers_premium[i]);j++{
			log.Println(i, " ",j," ", numbers_premium[i][j])
		}
	}
}
package main

import(
	"log"
	"slices"
)


func main(){
	numbers := []int{10,2,3,1,0,-1,5}

	slices.Sort(numbers)
	slices.Reverse(numbers)
	log.Println(numbers)
}
package main

import (
	"fmt"
	"log"
)

type Person struct{
		Name string
		Age int
	}
//methods 
func (p Person) GetName() string{
	return p.Name
}
func (p Person) GetAge() int{
	return p.Age
}

func main(){

	p:= Person{Name:"Alice",Age:50} 

	log.Println(p)
	fmt.Println("Name :",p.Name)
	fmt.Println("Age :",p.Age)
	
	
	
}
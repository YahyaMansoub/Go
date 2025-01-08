package main


import (
	"fmt"
)


type Person struct{
	Name string
	Age int
}

type Employee struct{
	Person
	EmployeeID interface
}

func main(){
	e:=Employee{}
	e.Name="Name"
//or
	e.Person.NAME="Name"
	e.Age=20
	e.EmployeeID=4512
//or
	p:= Employee{
		Person{
			Name:"Name",
			Age:20,
		},
		EmployeeID:15423,
	}
	ep := p.Person
	PersonDetails(ep)

}



func PersonDetails(p Person){
	fmt.Println("The Persons's details are : ",p.Name," ",p.Age)
}
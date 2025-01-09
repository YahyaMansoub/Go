package main


import (
	"log"
	"fmt"
	"strconv"
	"encoding/json"
	"os"
	"slices"
)

func avg(list []int)int{
	var avg int
	for _,value := range list{
		avg += value
	}
	avg/= len(list)
	return avg
}

type Person struct{
	name string
	age int
	salary int
	education string
}
type People struct{
	People []Person 
}


func main(){

	a := []int{1,4,8,5,1,2}
	average := avg(a)
	log.Println(average)

	File, err := os.ReadAll("./person_input.json")
	if err != nil{
		log.Fatal(err)
	}
    people:=People{}

    err=json.Unmarshal(File, &people)
	if err != nil{
		log.Fatal(err)
	}


    for i := 0; i < len(people.People); i++ {
          fmt.Println("Name : " + people.People[i].name)
          fmt.Println(" Age: " + strconv.Itoa(people.People[i].age) )
          fmt.Println(" Salary: " + strconv.Itoa(people.People[i].salary))
          fmt.Println(" Education: " + people.People[i].education)
}   
    var names,educations []string
	var ages,salaries []int
	
	

    for i := 0; i < len(people.People); i++ {
		  p:=people.People[i] 
		  youngest:=p.name

          names = append(names,p.name)
          ages = append(ages,p.age)
          salaries = append(salaries,p)
          educations = append(educations,p)
}
    var MapofPeople map[string]string
    AverageAge:=avg(ages)
	AverageSalary:=avg(salaries)
    for i := 0; i < len(people.People); i++ {
		  p:=people.People[i] 
		  if i>0{
			    if minage>p.age{
				     yougest=p.name
			    } 
				if maxage<p.age{
				     oldest=p.name
			    } 
				if maxsal<p.age{
				     oldest=p.name
			    } 
				if minsal<p.age{
				     oldest=p.name
			    } 
		  }else {
			youngest:=p.name
			minage:=p.age
			oldest:=p.name
			maxage:=p.age
			HighestSalary:=p.name
			maxsal:=p.salary
			LowestSalary:=p.name
			minsal:=p.salary
		  } 
		  MapofPeople[p.name] =p.education          
}
    
}
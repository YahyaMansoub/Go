package main


import (
	"log"
	"os"
	"encoding/json"
)

type Person struct{
	Name string
	Age int
}

func main(){
	file, err:=os.ReadFile("./person.json")
	if err != nil{
		log.Fatal(err)
	}
	p:= Person{}
	err = json.Unmarshal(file,&p)
	if err != nil{
		log.Fatal(err)
	}
	log.Println(p.Name, p.Age)
}
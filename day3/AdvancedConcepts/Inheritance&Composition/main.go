


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

	p:= Employee{
		Person{
			Name:"Name",
			Age:0,
		},
		EmployeeID:15423,
	}
}
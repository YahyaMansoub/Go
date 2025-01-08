package main


import (
	"fmt"
)

type Vehicule struct{
	Make string
	Model string 
	Year int
}

type Insurable interface{
	CalculateInsurance()
}
type Printable interface{
	Details()
}

type Car struct{
	Vehicule
	NumberofDoors int
}
func (c Car) CalculateInsurance() float64{
	return 0.98
}
func (c Car) Details() string{
	return "details for a car"
}
type Truck struct{
	Vehicule
	PayloadCapacity float64
}
func (c Truck) CalculateInsurance() float64{
	return 0.90
}
func (c Truck) Details() string{
	return "details for a Truck"
}


func main(){
	c1:=Car{
		Vehicule{
			Make:"Hyundai",
			Model:"Late",
			Year:2004,
		},
		NumberofDoors:4,
	}
	t1:=Truck{
		Vehicule{
			Make:"Hyundai",
			Model:"Late",
			Year:2004,
		},
		PayloadCapacity:4520,
	}

	fmt.Println(t1.Details())
	fmt.Println(c1.Details())
}



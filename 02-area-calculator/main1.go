package main
import "fmt"


func main(){
	var x float64= 2
	var y float64= 47

	var area float64= x*y 

	fmt.Printf("The area of the rectangle in meters is %.2f\n",area)
	const conv = 3.28
	var convert float64= conv*area
	fmt.Printf("The area of the rectangle in foot is %.2f",convert)



}
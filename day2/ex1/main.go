package main
import "fmt"

func rectangle( width,length float64) float64{
	return width * length
}
func rectarray(arr [2]float64) float64{
	return arr[0] * arr[1]
}

func main(){
	var a float64 = 5.0
	var b float64 = 42.65
	var arr [2]float64 = [2]float64{a,b}
	fmt.Println("Area of rectangle is ", rectangle(a, b))
	fmt.Println("Area of rectangle is ", rectarray(arr))

}
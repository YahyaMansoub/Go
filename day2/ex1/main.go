package main
import "fmt"

func rectangle( width,length float64) float64{
	return width * length
}
func rectarray(arr [2]float64) float64{
	return arr[0] * arr[1]
}
func rectareacheck(w,l float64) (float64,bool){
	if l<0 || w<0{
		return 0,false
	}
	return w*l,true
}	

func increment() func() int{
	i:=0
	return func() int{
		i++
		return i
	}
}

func main(){
	var a float64 = 5.0
	var b float64 = 42.65
	var arr [2]float64 = [2]float64{a,b}
	area := func(width float64,length float64)float64{
		return width * length
	}
	fmt.Println("Area of rectangle is ", rectangle(a, b))
	fmt.Println("Area of rectangle is ", rectarray(arr))
	fmt.Println("Area of rectangle is ", area(a,b))
	fmt.Println(rectareacheck(a,b))
	fmt.Println(increment())


}
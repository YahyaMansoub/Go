package main
import "fmt"

func sum (numbers ...int) int {
		total:=0
		for i:=0; i< len(numbers);i++{
			total += numbers[i]
		}
		return total
}

func rectangle(width, length float64)(area,perimiter float64) {
	area = width * length
	perimiter = 2 * (width + length)
	return
}
//Closures

func main(){
	fmt.Println(rectangle(4,2))
	//anonymous functions
	add:= func(a int,b int) int{
		return a+b
	}	
	// anonymous functions are only the one possibel to use in the main function, it is limited to the function in which it was declared.
	fmt.Println(add(5,6))
}
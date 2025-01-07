package main
import "fmt"

func sum (numbers ...int) int {
		total:=0
		for i:=0; i< len(numbers);i++{
			total += numbers[i]
		}
		return total
}

func rectangle(width, length float64)(area , perimiter float64) {
	area = width * length
	perimiter = 2 * (width + length)
	return
}
func main(){

	fmt.Println(rectangle(4,2))


	
}
package main
import (
	"fmt"
	"example.com/ex/areas"
	
)




func main(){
	var x,y float64  = 12, 47
	 area   := areas.Calculatearea(x,y)
	
	fmt.Println("Area ",area)
}
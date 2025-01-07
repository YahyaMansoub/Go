package main


import(
	"log"
)
func printNumbers(numbers []int){
		for index,value:= range numbers{
			log.Println(index,":",value)
		}
	}
func printNumberswithpointers(numbers []int){
		for index,value:= range numbers{
			log.Println("%v : %v : %p", index, value, )
		}
	}
func main(){
	
	
	a:=[]int{4,7,8,4}
	printNumbers(a)
}
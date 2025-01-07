package main


import "fmt"



func increment() func() int{
	i:=0
	return func()int{
		i++
		return i
	}
}


func main(){
	f:=increment()
	fmt.Println(f()) 
	fmt.Println(f())
	fmt.Println(f()) 

	

}
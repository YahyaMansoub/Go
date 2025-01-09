package main


import (
	"fmt"
)

func main(){
	v:=10

	ch:=make(chan int)

	go func(){
		s:=square(v)
		ch <- s
		ch <- 45
		fmt.Println("ehm waalaaa")
		ch <- 100
	}()

	vSquare := <- ch
	fmt.Println(v,"squared =",vSquare)
	vSqu := <- ch
	fmt.Println(vSqu)
	vSquar := <- ch
	fmt.Println(vSquar)


} 


func square(v int) int{
	return v*v
}




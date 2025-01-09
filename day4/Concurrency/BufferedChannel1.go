package main


import (
	"fmt"
)


func main(){

	ch:= make(chan int, 5)

	go func(){
		for i:= range 10{
			ch <- i
		}
		close(ch)
	}()
	for {
		v,ok:= <- ch
		if !ok{
			break
		}
		fmt.Println("received in main: ",v)
	}
}
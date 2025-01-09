package main


import (
	"fmt"
	"time"
)

func generatePrimes(n int, ch chan<- int){

	for i:= range n{
		ch <- i+1
	}
	close(ch)
}

func printPrimes(ch <-chan int){
	for v := range ch{
		fmt.Println(v)
	}
}
func main(){

	ch:= make(chan int)
	go generatePrimes(10,ch)
	go printPrimes(ch)
	time.Sleep(100*time.Microsecond)
}
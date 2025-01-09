package main


import (
	"fmt"
	"time"
)

var counter = 0
func increment(){
	counter++
}

func main(){
	for _ = range 1000{
		go increment()
	}
	time.Sleep(2*time.Second)
	fmt.Println(counter)
}
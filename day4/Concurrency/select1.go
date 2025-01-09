package main



import (
	"fmt"
	"time"
)

func main(){

	ch:=time.After(2*time.Second)
	loop:
	     for{
		select {
	case <- ch:
		fmt.Println("After 2 seconds")
		break loop
	default:
		fmt.Println("Still waiting")
	}
	}
}
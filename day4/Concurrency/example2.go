package main 


import (
	"fmt"
	"time"
)


func main(){
	message :="message"

	go func(){
		fmt.Println(message)
	}()
	message = "Hello World"
	message = "nice"
	time.Sleep(time.Millisecond)
}
package main


import(
	"fmt"
	"time"
)



func sayHello(){
	fmt.Println("Hello")
}


func main(){
	// Go routines 
	go sayHello()
	// or instead here time.Sleep(time.Second)
	fmt.Println("World")
	time.Sleep(time.Second)
	go do1()
	go do2()
	go do3()
	time.Sleep(8 *time.Second)
}

func do1(){
	time.Sleep(4 * time.Second)
	fmt.Println("1")
}
func do2(){
	time.Sleep(3 * time.Second)
	fmt.Println("2")
}
func do3(){
	time.Sleep(7 * time.Second)
	fmt.Println("3")
}
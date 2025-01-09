package main


import(
	"fmt"
	"context"
	"time"
)



func main(){
	// ctx := context.Background() // Graceful stop (without crushing the server )
	ctx , cancel := context.WithCancel()
	go doWork(ctx)
	time.Sleep(5 * time.Second)
	defer cancel()
}


func doWork(ctx context.Context){
	time.Sleep(time.Second)
	fmt.Println("Hello")
	for{
		select{
	case <- ctx.Done():
		fmt.Println("Stopping Gracefully")
		return
	default:
		fmt.Println("still here")
	}
}
}
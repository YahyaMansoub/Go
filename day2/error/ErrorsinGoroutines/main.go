package main

import (
	"fmt"
	"time"
)

func processJob(job string, ch chan error) {
	if job == "bad" {
		ch <- fmt.Errorf("error processing job %s", job)
		return
	}
	ch <- nil
}

func main() {
	ch := make(chan error)

	go processJob("good", ch)
	go processJob("bad", ch)

	for i := 0; i < 2; i++ {
		err := <-ch
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Job processed successfully")
		}
	}
}

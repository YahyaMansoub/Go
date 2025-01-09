package main


import (
	"fmt"
	"time"
	"context"
	"net/http"
	"sync"
)


func main(){

	fmt.Println("initializing Scrapping... ")

	urls := []string{"http://example.com", "http://httpbin.org/delay/5"}
	timeout:= 3* time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var wg sync.WaitGroup

	for _,url := range urls{
		wg.Add(1)
		go fetchURL(ctx,url,&wg)
	}

	wg.Wait()

	fmt.Println("Scraping complete")
}

func fetchURL(ctx context.Context, url string , wg *sync.WaitGroup){
	defer wg.Done()

	req, err:= http.NewRequest("GET", url, nil)
	if err != nil{
		
		fmt.Println("Error creating request for",url,":",err)
	}
	req = req.WithContext(ctx)

	client:= &http.Client{}
	resp, err := client.Do(req)
	if err!=nil{
		if ctx.Err()!=nil{
			fmt.Println("Request for ",url," was cancelled/timed out:",ctx.Err())
	
		}else{
			fmt.Println("Error fetching ",url,":", err)
		}
		return
	}
	defer resp.Body.Close()

	fmt.Println("Successfully fetched ", url)


}
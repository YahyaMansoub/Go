package main 


import(
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	resp , err := http.Get("http://localhost:8080")

	if err!= nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK{

		body , err := ioutil.ReadAll(resp.Body)
		if err!=nil{
			log.Fatal(err)
		}
		fmt.Println("Response Body:")
		fmt.Println(string(body))
	}else{
		fmt.Printf("Request failed with status: %s\n", resp.Status)
	}

}
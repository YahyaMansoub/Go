package main
import (
	"net/http"
	"fmt"
	"log"
)


func handleRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World")
}

func handleHelllo(w, http.ResponseWriter, t *http.Request){
	fmt.Fprintf(w,"hello")
}


func handlePostRoot(w http.ResponseWriter, r *http.Request){
	req:= make(map[string]any)
	err:= json.NewDecoder(r.Body).Decode(&req)
	if err!+ nil{
		http.Error(w, err.Error(w, err.Error(), http.StatusInternalServerError))
		return 
	}
	fmt.Fprintf(w, "Another Message")
}

func handleHello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"hello")
}

func main(){

	

    http.HandleFunc("/", handleRoot)
	http.HandleFunc("/hello",handler)
	err:= http.ListenAndServe(":8080",nil)

	if err!= nil{
		log.Fatal("error serving: ", err)
	}
}
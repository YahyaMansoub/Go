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

func main(){

	paths := []string{"path1","path2","path3"}
	for _, path := range paths {
    http.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, path)
    })
}

	http.HandleFunc("/hello",handler)
	err:= http.ListenAndServe(":8080",nil)

	if err!= nil{
		log.Fatal("error serving: ", err)
	}
}
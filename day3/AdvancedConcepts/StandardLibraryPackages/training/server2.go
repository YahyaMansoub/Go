package main



import (
	"fmt"
	"net/http"
)


func Handler(m string,w http.ResponseWriter ,r *http.ResponseWriter ){
	fmt.Fprintf(w,m)
}
func main(){

	response:=[]string{"hallo","bruh","ehm","wellwell"}

	for _,value := range response{
		http.HandleFunc("/"+value, Handler(value))
	}
	http.ListenAndServe(":8080",nil)

}
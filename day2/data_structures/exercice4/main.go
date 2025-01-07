package main


import (
	"log"
)
func freq(S []string)map[string]int{
	Map := make(map[string]int)
	for _, value := range S{
		_,ok := Map[value]
		if ok{
			Map[value]++
		}else{
			Map[value]=1
		}
	}
	return Map

}

func main(){
	slice:= []string{"aba","h","h","hg","aba","hg"}
	Map := freq(slice)

	log.Println(Map)
	

}
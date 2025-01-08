package main 
import(
	"os"
)


type JSONFile struct{
	People []Person
}

type Person struct{
	Name string
	Age int
	Salary int
	Education string
}


func main(){



	file,err := os.ReadFile("./people.json")
	if err!=nil{
		log.Fail("cou",err)
	}

	f:=JSONFile{}
	// f:=struct {People []Person}{}

	err=json.Unmarshal(file, f)
	if err!=nil{
		log.Fatal("could",err)
	}
}


func GetAverageAge(p []Person)(int,err){
	sum:=0
	if len(p)==0{
		return 0,fmt.Errorf("impossible")
	}
	for _,v := range p{
		sum+=v.Age
	}
	return sum/len(p), nil
}


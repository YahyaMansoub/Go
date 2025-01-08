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

type Output struct{
	AvrafeAge int  
}


func main(){



	file,err := os.ReadFile("./people.json")
	if err!=nil{
		log.Fail("cou",err)
	}

	f:=JSONFile{}
	// f:=struct {People []Person}{}

	err=json.Unmarshal(file,&f)
	if err!=nil{
		log.Fatal("could",err)
	}

	o:=Output{
		Average: averageage,
	}

	outputFile, err:=json.Marshal(o)
	if err!= nil{
		log.Fatal("",err)

		
		
	}
	if err:= os.WriteFile("./ouput.json",outputFile){
		log.Fatal("",err)
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


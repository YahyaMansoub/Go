package main
import (
    "net/http"

    "github.com/gin-gonic/gin"
)
func deleteElement(slice []student, index int) []student {
   return append(slice[:index], slice[index+1:]...)
}

type student struct {
    ID     string  `json:"id"`
    Name  string  `json:"name"`
    Age int  `json:"age"`
    Address  string `json:"address"`
    Courses  string `json:"courses"`
}


var students = []student{
    {ID: "1", Name: "Abdellah", Age: 26, Address :"derb deban ",Courses:" No Courses"},
    {ID: "2", Name: "El-Hachmi", Age: 45, Address :"derb deban ", Courses:"No Courses"},
}
func getStudents(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, students)
}

func postStudents(c *gin.Context) {
    var newStudent student 

    if err := c.BindJSON(&newStudent); err != nil {
        return
    }
    students = append(students, newStudent)
    c.IndentedJSON(http.StatusCreated, newStudent)
}

func putStudentByID(c *gin.Context){
	id := c.Param("id")
	var updatedStudent student
	if err := c.BindJSON(&updatedStudent); err != nil {
        return
    }
	for i,a := range students{
		if a.ID == id {
			students[i]=updatedStudent
			c.IndentedJSON(http.StatusUpdated, updatedStudent)
			return 
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})

}

func deleteStudentByID(c *gin.Context){
	id := c.Param("id")

	for i,a := range students{
		if a.ID == id {
			students =deleteElement(students,id)
			c.IndentedJSON(http.StatusOK,a)
			return 
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}


func getStudentByID(c *gin.Context){
	id := c.Param("id")
	for _,a := range students{
		if a.ID == id {
			c.IndentedJSON(http.StatusOK,a)
			return 
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}
func getStudentByNAME(c *gin.Context){
	name := c.Param("name")
	for _,a := range students{
		if a.Name == name {
			c.IndentedJSON(http.StatusOK,a)
			return 
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}

func main() {
    router := gin.Default()
    router.GET("/students", getStudents)
    router.GET("/students/:id", getStudentByID)
    router.GET("/students/:name", getStudentByNAME)
	router.POST("/students",postStudents)
    router.PUT("/students/:id", putStudentByID)
    router.DELETE("/students/:id", deleteElement)
    router.Run("localhost:8080")
}
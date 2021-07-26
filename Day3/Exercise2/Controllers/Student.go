package Controllers


import (
	"awesomeProject/Day3/Exercise2/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
//Getstudents ... Get all students
func GetStudents(c *gin.Context) {
	var student []Models.Student
	err := Models.GetAllStudents(&student)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}
//Createstudent ... Create student
func CreateStudent(c *gin.Context) {
	var student Models.Student
	c.BindJSON(&student)
	err := Models.CreateStudent(&student)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}
//GetstudentByID ... Get the student by id
func GetStudentByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var student Models.Student
	err := Models.GetStudentByID(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}
//Updatestudent ... Update the student information
func UpdateStudent(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")
	err := Models.GetStudentByID(&student, id)
	if err != nil {
		c.JSON(http.StatusNotFound, student)
	}
	c.BindJSON(&student)
	err = Models.UpdateStudent(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}
//Deletestudent ... Delete the student
func DeleteStudent(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")
	err := Models.DeleteStudent(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

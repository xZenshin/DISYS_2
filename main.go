package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
)

type Course struct {
	Name string `json:"name,omitempty"`

	CourseID int `json:"courseID,omitempty"`

	Workload int32 `json:"workload,omitempty"`

	StudentSatisfaction int32 `json:"studentSatisfaction,omitempty"`
}

func main() {
	log.Printf("Server started")

	router := gin.Default()
	router.GET("/courses", getCourses)
	router.POST("/course", postCourse)
	router.GET("/course/:id", getCourse)
	router.POST("/updateCourse", upDeleteCourse)
	router.DELETE("/courseDel/:id", deleteCourse)

	router.Run("localhost:8080")
}

var courses = []Course{
// {Name: "Benny", CourseID: 2, Workload: 1, StudentSatisfaction: 1},
}

func postCourse(c *gin.Context) {
	var newCourse Course

	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusCreated, newCourse)
}

func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

func getCourse(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))
	for _, a := range courses {
		if a.CourseID == id {
			c.IndentedJSON(http.StatusOK, a)
		}
	}
}

func deleteCourse(c *gin.Context) {
	var index = 0
	var id, _ = strconv.Atoi(c.Param("id"))
	for _, a := range courses {

		if a.CourseID == id {
			courses = RemoveIndex(courses, index)
			c.IndentedJSON(http.StatusOK, a)
		}
		index++
	}
}
func RemoveIndex(s []Course, index int) []Course {
	return append(s[:index], s[index+1:]...)
}

func upDeleteCourse(c *gin.Context){
	var courseToDelete Course
	if err := c.BindJSON(&courseToDelete); err != nil {
		return
	}
	var index = 0
	for _, a := range courses {

		if a.CourseID == courseToDelete.CourseID {
			courses = RemoveIndex(courses, index)
			c.IndentedJSON(http.StatusOK, a)
		}
		index++
	}

	courses = append(courses, courseToDelete)
	c.IndentedJSON(http.StatusCreated, courseToDelete)

}

func updateCourse(c *gin.Context) {
	var updateCourse Course
	if err := c.BindJSON(&updateCourse); err != nil {
		return
	}
	
	for _, a := range courses {
		if updateCourse.CourseID == a.CourseID {
			p := &a
			p.Name = updateCourse.Name
			fmt.Println(a.Name)
			p.StudentSatisfaction = updateCourse.StudentSatisfaction
			p.Workload = updateCourse.Workload
			c.IndentedJSON(http.StatusOK, a)
		}
	}

	
  }

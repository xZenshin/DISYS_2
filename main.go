package main

import (
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
	router.GET("/courseDel/:id", deleteCourse)

	router.Run("localhost:8080")
}

var courses = []Course{}

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

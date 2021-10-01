package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	sw "github.com/xZenshin/DISYS_2/go"
)

func main() {
	log.Printf("Server started")

	router := gin.Default()
	router.GET("/students", getStudents)
	router.POST("/students", postStudent)

	router.Run("localhost:8080")
}

var students = []sw.Student{}
var courses = []sw.Course{}

func postStudent(c *gin.Context) {
	var newStudent sw.Student

	if err := c.BindJSON(&newStudent); err != nil {
		return
	}

	students = append(students, newStudent)
	c.IndentedJSON(http.StatusCreated, newStudent)
}

func getStudents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, students)
}

func postCourse(c *gin.Context) {
	var newCourse sw.Course

	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusCreated, newCourse)
}

func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

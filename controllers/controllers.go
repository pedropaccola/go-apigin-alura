package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropaccola/go-apigin-alura/database"
	"github.com/pedropaccola/go-apigin-alura/models"
)

func GetAllStudents(c *gin.Context) {
	students := models.NewStudents()
	c.JSON(200, students)
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func Greetings(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API": "Greetings " + name + " !",
	})
}

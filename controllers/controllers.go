package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pedropaccola/go-apigin-alura/models"
)

func GetAllStudents(c *gin.Context) {
	students := models.NewStudents()
	c.JSON(200, students)
}

func Greetings(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API": "Greetings " + name + " !",
	})
}

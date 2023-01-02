package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/pedropaccola/go-apigin-alura/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/students", controllers.GetAllStudents)
	r.Run()
}

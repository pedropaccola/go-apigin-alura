package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/pedropaccola/go-apigin-alura/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.GetAllStudents)
	r.POST("/students", controllers.CreateStudent)
	r.GET("/students/:id", controllers.GetStudentById)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.GET("/:name", controllers.Greetings)
	r.Run()
}

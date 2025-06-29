package main

import (
	"hola/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Run(":8080") // API expuesta en http://localhost:8080
}

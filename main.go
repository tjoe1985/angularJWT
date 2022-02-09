package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tjoe1985/angularJWT.git/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/", controllers.Home())
	//	r.POST("/register", controllers.Register)
	r.Run(":8080")
}

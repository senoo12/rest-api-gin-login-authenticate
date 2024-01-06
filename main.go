package main

import (
	"github.com/gin-gonic/gin"
	"ProductAPI/controllers"
	"ProductAPI/models"
)

func main()  {
	r := gin.Default()
	models.ConnectDatabse()

	r.POST("/api/register", controllers.Register)
	r.POST("/api/login", controllers.Login)

	r.Run()
}
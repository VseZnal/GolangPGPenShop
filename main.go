package main

import (
	"GolangPGPenShop/controllers"
	"GolangPGPenShop/models"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.Use(cors.Default())

	r.GET("/items", controllers.FindItems)
	r.POST("/items", controllers.CreateItemNew)
	r.GET("/items/:id", controllers.FindItem)

	r.Run()
}

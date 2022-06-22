package main

import (
	"GolangPGPenShop/controllers"
	"GolangPGPenShop/models"
	"net/http"

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
	r.GET("/admin", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.youtube.com/watch?v=dQw4w9WgXcQ&ab_channel=RickAstley")
	})

	r.Run()
}

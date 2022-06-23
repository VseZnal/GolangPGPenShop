package main

import (
	"GolangPGPenShop/controllers"
	"GolangPGPenShop/middlewares"
	"GolangPGPenShop/models"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.Use(cors.Default())

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	public.GET("/items", controllers.FindItems)
	public.GET("/items/:id", controllers.FindItem)
	public.GET("/admin", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently,
			"https://www.youtube.com/watch?v=dQw4w9WgXcQ&ab_channel=RickAstley")
	})

	protected := r.Group("/api/ad")
	protected.Use(middlewares.JwtAuthMiddleware())

	protected.PATCH("/items/:id", controllers.UpdateItemNew)
	protected.DELETE("/items/:id", controllers.DeleteItem)
	protected.POST("/items", controllers.CreateItemNew)
	protected.GET("/user", controllers.CurrentUser)

	r.Run()
}

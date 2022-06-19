package main

import (
	"GolangPGPenShop/models"
	"github.com/astaxie/beego/orm"
)

var ORM orm.Ormer

func init() {
	models.ConnectToDb()
	ORM = models.GetOrmObject()
}

/*
func main() {
	router := gin.Default()
	router.POST("/createUser", createUser)
	router.GET("/readUsers", readUsers)
	router.PUT("/updateUser", updateUser)
	router.DELETE("/deleteUser", deleteUser)
	router.Run(":8000")
}
*/

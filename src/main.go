package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ryonryon/Go-Practice/src/controllers"
	"github.com/ryonryon/Go-Practice/src/interfaces/handler"
	"github.com/ryonryon/Go-Practice/src/models"
)

var (
	R *gin.Engine
)

func init() {
	R = gin.Default()
	handler.DB.AutoMigrate(models.User{})
	fmt.Print("[RYON-debug] Migrated\n")
	R.GET("/users", controllers.UserController.IndexUserController)
	R.GET("/users/:id", controllers.UserController.IndexOneUserController)
	R.POST("/users/create", controllers.UserController.CreateUserController)
	R.PUT("/users/:id", controllers.UserController.UpdateUserController)
	R.DELETE("/users/:id", controllers.UserController.DestroyUserController)
}

func main() {
	err := R.Run() // デフォルトで:8080になる
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ryonryon/Go-Practice/src/controllers"
	"github.com/ryonryon/Go-Practice/src/interfaces/handler"
	"github.com/ryonryon/Go-Practice/src/models"
)

func main() {
	r := gin.Default()
	db := handler.CreateConnection()
	db.AutoMigrate(models.User{})
	fmt.Print("[RYON-debug] Migrated\n")
	// r.GET("/hello", controllers.GetHello) // 関数の頭文字が大文字じゃないと、参照できない
	r.GET("/users", controllers.IndexUserController)
	r.GET("/users/:id", controllers.IndexOneUserController)
	r.POST("/users/create", controllers.CreateUserController)
	r.DELETE("/users/:id", controllers.DestroyUserController)
	r.Run() // デフォルトで:8080になる
}

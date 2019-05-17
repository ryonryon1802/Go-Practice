package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ryonryon/Go-Practice/src/controllers"
	"github.com/ryonryon/Go-Practice/src/interfaces/handler"
	"github.com/ryonryon/Go-Practice/src/models"
)

func main() {
	r := gin.Default()
	db := handler.CreateConnection()
	db.AutoMigrate(models.User{})
	// r.GET("/hello", controllers.GetHello) // 関数の頭文字が大文字じゃないと、参照できない
	r.GET("/user", controllers.GetUserController)
	r.Run() // デフォルトで:8080になる
}

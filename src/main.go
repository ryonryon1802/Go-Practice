package main

import (
	"./controllers"
	"./migrations"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := migrations.CreateConnection()
    defer db.Close()
    // r.GET("/hello", controllers.GetHello) // 関数の頭文字が大文字じゃないと、参照できない
	r.GET("/user", controllers.GetUser)
	r.Run() // デフォルトで:8080になる
}

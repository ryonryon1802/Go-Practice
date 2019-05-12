package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ryonryon/Go-Practice/src/controllers"
)

func main() {
	r := gin.Default()
    // r.GET("/hello", controllers.GetHello) // 関数の頭文字が大文字じゃないと、参照できない
	r.GET("/user?limit", controllers.GetUser)
	r.Run() // デフォルトで:8080になる
}

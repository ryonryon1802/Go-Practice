package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryonryon/Go-Practice/src/models"
)

// func GetHello(c *gin.Context) {
// 	c.String(200, "Hello,World!")
// }


func GetUser(c *gin.Context) {
	user := user_models.GetUsers()
	c.JSON(200, user)
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryonryon/Go-Practice/src/models"
)

// func GetHello(c *gin.Context) {
// 	c.String(200, "Hello,World!")
// }

func GetUserController(c *gin.Context) {
	user := models.User{}
	users := user.GetUser()

	c.JSON(200, users)
}
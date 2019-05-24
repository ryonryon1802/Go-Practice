package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryonryon/Go-Practice/src/models"
	"log"
	"net/http"
)

// func GetHello(c *gin.Context) {
// 	c.String(200, "Hello,World!")
// }

func GetUserController(c *gin.Context) {
	user := models.User{}
	users := user.GetUser()
	c.JSON(http.StatusOK, users)
}

func GetOneUserController(c *gin.Context) {
	user := models.User{}
	id := c.Param("id")
	users := user.GetOneUser(id)
	c.JSON(http.StatusOK, users)
}

func AddUserController(c *gin.Context) {
	user := new(models.User)
	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err})
		log.Print(err)
		return
	}
	user.AddUser(user)
}
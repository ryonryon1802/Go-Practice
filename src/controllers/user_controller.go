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

func IndexUserController(c *gin.Context) {
	user := models.User{}
	users := user.IndexUser()
	c.JSON(http.StatusOK, users)
}

func IndexOneUserController(c *gin.Context) {
	user := models.User{}
	id := c.Param("id")
	users, err := user.IndexOneUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": err.Error()})
		log.Print(err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func CreateUserController(c *gin.Context) {
	user := new(models.User)
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err})
		log.Print(err)
		return
	}
	user.CreateUser(user)
}
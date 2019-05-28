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
	users, err := user.IndexUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		log.Print(err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func IndexOneUserController(c *gin.Context) {
	user := models.User{}
	id := c.Param("id")
	users, err := user.IndexOneUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		log.Print(err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func CreateUserController(c *gin.Context) {
	user := &models.User{}
	err := user.CreateUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		log.Print(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func UpdateUserController(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}
	err := user.UpdateUser(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		log.Print(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func DestroyUserController(c *gin.Context) {
	user := models.User{}
	id := c.Param("id")
	err := user.DestroyUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		log.Print(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
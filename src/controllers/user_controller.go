package controllers

import (
	"github.com/gin-gonic/gin"
	"../models"
	"../migrations"
)

// func GetHello(c *gin.Context) {
// 	c.String(200, "Hello,World!")
// }

// func GetUser(c *gin.Context) {
// 	db := migrations.CreateConnection()

// 	users := []models.User{}
// 	db.Find(&users)
// 	c.JSON(200, users)
// }
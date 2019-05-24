package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ryonryon/Go-Practice/src/models"
	"reflect"
)

// func GetHello(c *gin.Context) {
// 	c.String(200, "Hello,World!")
// }

func GetUserController(c *gin.Context) {
	user := models.User{}
	users := user.GetUser()

	// debug
	fmt.Printf("Controller: %p\n", users)
	fmt.Printf("Controller: %v\n", users)
	fmt.Printf("Controller: %v\n", *users)
	fmt.Println(reflect.TypeOf(users))

	c.JSON(200, users)
}

func AddUserController(c *gin.Context) {
	user := new(models.User)

	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
	}

	user.AddUser(user)
}
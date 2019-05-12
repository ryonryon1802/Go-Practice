package models

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ryonryon/Go-Practice/src/interfaces/handler"
	"time"
)

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) GetUsers(c *gin.Context) {
	db := gorm_handler.CreateConnection()
	// limit := c.Param("limit")

	// user := new(User) //単体
	users := []user{}
	db.Find(users)
	c.JSON(200, user)
}

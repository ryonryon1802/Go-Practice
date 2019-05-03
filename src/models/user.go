package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) getUser(c *gin.Context) {
	db := migrations.CreateConnection()

	user := new(User)
	db.Find(user)
	c.JSON(200, user)
}

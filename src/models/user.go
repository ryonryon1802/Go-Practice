package models

import (
	"github.com/jinzhu/gorm"
	"github.com/ryonryon/Go-Practice/src/interfaces/handler"
	"time"
)

// gorm.Modelの定義
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt *time.Time
}

type User struct {
	gorm.Model
	Name string `json:"name"`
}

//func InitializeUser(db *gorm.DB) {
//	db.AutoMigrate(User{})
//}

func (u *User) GetUser() *[]User {
	db := handler.CreateConnection()
	user := &[]User{}
	db.Find(user)
	return user
}

func (u *User) GetOneUser(id string) *[]User {
	db := handler.CreateConnection()
	user := &[]User{}
	db.Where("id = ?", id).Find(user)
	return user
}

func (u *User) AddUser(user *User) {
	db := handler.CreateConnection()
	db.NewRecord(&user)
	db.Create(&user)
}
package models

import (
	"fmt"
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

func (u *User) IndexUser() *[]User {
	db := handler.CreateConnection()
	user := &[]User{}
	db.Find(user)
	return user
}

func (u *User) IndexOneUser(id string) (*[]User, error) {
	db := handler.CreateConnection()
	user := &[]User{}
	db.Where("id = ?", id).Find(user)
	if id == "" {
		return nil, fmt.Errorf("parameter cannot find")
	}
	if cap(*user) == 0 {
		return nil, fmt.Errorf("cannot find id: %s", id)
	}
	return user, nil
}

func (u *User) CreateUser(user *User) {
	db := handler.CreateConnection()
	db.NewRecord(&user)
	db.Create(&user)
}
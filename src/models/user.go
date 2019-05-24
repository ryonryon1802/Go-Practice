package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/ryonryon/Go-Practice/src/interfaces/handler"
	"reflect"
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

func (u *User) GetUser() *[]User { // *[]Userはポインタ変数？
	db := handler.CreateConnection()

	user := &[]User{}

	// debug
	fmt.Printf("Model: %p\n", user)
	fmt.Printf("Model: %v\n", user)
	fmt.Printf("Model: %v\n", *user)
	fmt.Println(reflect.TypeOf(user))

	db.Find(user)

	return user
}

func (u *User) AddUser(user *User) {
	db := handler.CreateConnection()

	db.NewRecord(&user)
	db.Create(&user)
}
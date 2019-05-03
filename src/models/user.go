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


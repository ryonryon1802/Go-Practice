package interfaces

import (
    "github.com/jinzhu/gorm"
    "fmt"
)

func CreateConnection() *gorm.DB {
    DBMS     := "mysql"
    USER     := "root"
    PASS     := "pass"
    PROTOCOL := "tcp(192.168.33.10:3306)"
    DBNAME   := "test"
    CONNECT  := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?parseTime=True&loc=Local"
    
    DB, err := gorm.Open(DBMS, CONNECT)

    if err != nil {
        panic(err) // panicは、どうしようもないエラーのときに利用 強制的にプログラムを落とす
    }
    fmt.Printf("[RYON-debug] Connected %s\n", DBMS)
    // DB.AutoMigrate(&models.User{})
    return DB
}

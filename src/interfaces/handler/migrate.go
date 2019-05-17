package handler

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
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
        panic(err) // panicは、どうしようもないエラーのときに利用 無理やり、プログラムを落とす
    }
    fmt.Printf("[RYON-debug] Connected %s\n", DBMS)
    return DB
}

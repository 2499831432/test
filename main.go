package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
//用户名：密码^@tcp(地址:3306)/数据库
db, err := sql.Open("mysql", "root:123456%^@tcp(39.107.87.114:3306)/book_manager?charset=utf8")
if err!=nil {
fmt.Println(err)
return
}

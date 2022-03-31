package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"	// 导入但不使用
)

type user struct {
	Id int 
	Name string
	Age int
}

var db *sqlx.DB

func initDB() (err error){
	// 链接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/study"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil{
		return
	}
	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return
}

func sqlInject(name string){
	sql_string := fmt.Sprintf("select id, name, age from user where name='%s'", name)

	fmt.Printf("SQL: %s\n", sql_string)

	var users []user
	err := db.Select(&users, sql_string)
	if err != nil{
		fmt.Println("get error: ", err)
		return
	}
	for _, u := range users{
		fmt.Printf("user: %#v\n", u)
	} 
}

func main() {
	err := initDB()
	if err != nil{
		fmt.Println("init error:", err)
	}
	// sql注入的几种示例
	// sqlInject("hunt")
	sqlInject("xxx' or 1=1 #")
}
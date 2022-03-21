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

func main() {
	err := initDB()
	if err != nil{
		fmt.Println("init error:", err)
	}
	
	sql_string1 := "select id, name, age from user where id=1"
	var u user
	err = db.Get(&u, sql_string1)
	if err != nil{
		fmt.Println("get failed error:", err)
	}
	fmt.Printf("u: %#v\n", u)

	var userlist []user
	sql_string2 := "select id, name, age from user"
	db.Select(&userlist, sql_string2)
	fmt.Printf("userlist: %#v\n", userlist)
}
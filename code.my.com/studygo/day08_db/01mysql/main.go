package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"	// 导入但不使用
)

type user struct {
	id int
	name string
	age int
}

var db *sql.DB

func initDB() (err error){
	// 链接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/study"
	db, err = sql.Open("mysql", dsn)	// 不会校验用户名和密码
	if err != nil{	// dsn格式不正确的时候会报错
		fmt.Printf("open %s failed. err: %v\n", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil{
		return
	}
	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return
}

func queryOne(id int) {
	var u1 user
	// 1. 写查询语句
	sql_string := `select id, name, age from user where id=?;`
	/*
	// 2. 执行
	row := db.QueryRow(sql_string, 1)
	for i := 1; i < 11; i++{
		fmt.Printf("The time %d query\n", i)
		db.QueryRow(sql_string, 1)
	}
	// 3. 结果
	row.Scan(&u1.id, &u1.name, &u1.age)
	*/

	// 必须对row对象使用scan方法，释放数据库链接
	db.QueryRow(sql_string, id).Scan(&u1.id, &u1.name, &u1.age)
	fmt.Printf("u1: %#v\n", u1)
}


func queryMore(n int) {
	// 1. 写查询语句
	sql_string := `select id, name, age from user where id>?;`
	rows, err := db.Query(sql_string, n)
	if err != nil{
		fmt.Printf("query %s err: %v\n", sql_string, err)
		return
	}
	// 一定要关闭每一个row对象
	defer rows.Close()
	for rows.Next(){
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil{
			fmt.Printf("scan err: %v\n", err)
			return
		}
		fmt.Printf("u1: %#v\n", u1)
	}
}

func insert() {
	sql_string := `insert into user(name, age) values("Smallmu", 24);`
	ret, err := db.Exec(sql_string)
	if err != nil{
		fmt.Printf("insert %s err: %v\n", sql_string, err)
		return
	}
	id, err := ret.LastInsertId()
	if err != nil{
		fmt.Printf("get id failed: %v\n", err)
		return
	}
	fmt.Println("id:", id)
}

func updateRow(newAge int, id int) {
	sql_string := `update user set age=? where id=?;`
	ret, err := db.Exec(sql_string, newAge, id)
	if err != nil{
		fmt.Printf("update %s err: %v\n", sql_string, err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil{
		fmt.Printf("get rows failed: %v\n", err)
		return
	}
	fmt.Printf("update %d rows\n", n)
}

func deleteRow(id int){
	sql_string := `delete from user where id=?;`
	ret, err := db.Exec(sql_string, id)
	if err != nil{
		fmt.Printf("delete %s err: %v\n", sql_string, err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil{
		fmt.Printf("get rows failed: %v\n", err)
		return
	}
	fmt.Printf("delete %d rows\n", n)
}

func main() {
	err := initDB()
	if err != nil{
		fmt.Printf("init database failed, err: %v\n", err)
	}
	fmt.Println("link success!")

	// queryOne(3)
	// insert()
	queryMore(0)
	// updateRow(25, 4)
}
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

func prepareSelct(id int) {
	var u1 user
	sqlStr := "select id, name, age from user where id=?;"
	//预处理SQL语句
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	//执行SQL，添加占位值
	str, err := stmt.Query(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	for str.Next() {
		err = str.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(u1.name)
}

// 与处理方式插入多条数据
func prepareInsert(m  map[string]int) {
	sqlStr := `insert into user(name, age) values(?, ?);`
	//预处理SQL语句
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	//执行SQL，填加站位值
	for k, v := range m{
		_, err = stmt.Exec(k, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	err := initDB()
	if err != nil{
		fmt.Printf("init database failed, err: %v\n", err)
	}
	fmt.Println("link success!")

	// prepareSelct(1)

	var m = map[string]int {
		"sj": 17,
		"lry": 16,
	}
	prepareInsert(m)
}
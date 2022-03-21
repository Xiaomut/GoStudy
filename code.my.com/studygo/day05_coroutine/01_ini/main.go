package main

import (
	"fmt"
)

type MysqlConfig struct{
	Address string `ini: "address"`
	Port int `ini: "port"`
	Username string `ini: "username"`
	Password string `ini: "password"`
}

type RedisConfig struct {
	Host string `ini: "host"`
	Port string `ini: "port"`
	Password string `ini: "password"`
	Database string `ini: "database"`
}

func loadini(file string, data interface{})  {
	// 1. 读ini
	// 2. 取数据	
}

func main()  {
	// var mc MysqlConfig
	// loadini(&mc)
	// fmt.Println(mc.Address, mc.Port, mc.Username, mc.Password)
}
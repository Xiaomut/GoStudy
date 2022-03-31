package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func GetRedisClient(){
    redisdb = redis.NewClient(&redis.Options{
        Addr:     "127.0.0.1:6379",
        Password: "", // no password set
        DB:       0,                 // use default DB
    })

    pong, err := redisdb.Ping().Result()
    if err != nil {
        fmt.Println(pong, err)
    }
}


func main() {
	GetRedisClient()
	fmt.Println("link redis success", redisdb)
}
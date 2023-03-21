package database

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// 声明一个全局的rdb变量
var Rdb *redis.Client

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "120.79.23.205:6379",
		Password: "485969746wqs", // no password set
		DB:       0,              // use default DB
		PoolSize: 100,            // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("连接Redis成功")
}

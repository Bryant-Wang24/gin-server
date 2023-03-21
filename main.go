package main

import (
	"fmt"

	"example.com/blog/database"
	"example.com/blog/routers"
)

func main() {
	database.InitMySQL()
	database.InitRedis()
	routers.SetupRouter()
	fmt.Println("Hello World")
}

package main

import (
	"example.com/blog/database"
	"example.com/blog/routers"
	"fmt"
)

func main() {
	database.InitMySQL()
	routers.SetupRouter()
	fmt.Println("Hello World")
}

package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func InitMySQL() {
	var err error
	dsn := "root:485969746wqs@tcp(localhost:3306)/gin_blog?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败", err)
		//panic("failed to connect database")
		return
	}
	fmt.Println("连接MySQL成功")
}

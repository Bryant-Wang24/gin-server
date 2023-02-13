package utils

import (
	"example.com/blog/database"
	"example.com/blog/model"
)

// CheckAuth 检查用户是否登录
func CheckAuth(username, password string) bool {
	var admin model.Admin
	database.Db.Where("username = ?", username).First(&admin)
	//如果用户名存在，且密码正确，返回true
	if admin.ID != 0 && admin.Password == password {
		return true
	}
	return false
}

// AddUser 添加用户
func AddUser(username, password string) bool {
	var admin model.Admin
	admin.Username = username
	admin.Password = password
	// 判断用户是否存在
	database.Db.Where("username = ?", username).First(&admin)
	// 如果用户存在，返回true;如果用户不存在，添加用户
	if admin.ID != 0 {
		return true
	} else {
		database.Db.Create(&admin)
		return false
	}
	// database.Db.Create(&admin)
	// if admin.ID != 0 {
	// 	return true
	// }
	// return false
}

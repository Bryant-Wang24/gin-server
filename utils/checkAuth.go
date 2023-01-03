package utils

import (
	"example.com/blog/database"
	"example.com/blog/model"
)

// CheckAuth 检查用户是否登录
func CheckAuth(userName, password string) bool {
	var admin model.Admin
	database.Db.Where("userName = ?", userName).First(&admin)
	//如果用户名存在，且密码正确，返回true
	if admin.ID != 0 && admin.Password == password {
		return true
	}
	return false
}

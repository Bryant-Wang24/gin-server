package utils

import (
	"example.com/blog/database"
	"example.com/blog/model"
)

// CheckAuth 检查用户是否登录,并返回用户ID
func CheckAuth(username, password string) (bool, int) {
	var admin model.Admin
	database.Db.Where("username = ?", username).First(&admin)
	//如果用户名存在，且密码正确，返回true
	if admin.ID != 0 && admin.Password == password {
		return true, admin.ID
	}
	return false, 0
}

// CheckAdmin 检查用户是否是管理员
func CheckAdmin(username string) bool {
	var admin model.Admin
	database.Db.Where("username = ?", username).First(&admin)
	//如果用户名存在，且是管理员，返回true
	if admin.ID != 0 && admin.IsAdmin == 1 {
		return true
	}
	return false
}

// AddUser 添加用户
func AddUser(username, password, from string) bool {
	var admin model.Admin
	admin.Username = username
	admin.Password = password
	// 判断用户是否存在
	database.Db.Where("username = ?", username).First(&admin)
	// 如果用户存在，返回true;如果用户不存在，添加用户
	if admin.ID != 0 {
		if from == "github" {
			// 如果是github登录，那么就更新密码
			database.Db.Model(&admin).Update("password", password)
			return false
		} else {
			return true
		}
	} else {
		database.Db.Create(&admin)
		return false
	}
}

// ChangePassword 修改密码
func CheckPassword(username, password string) bool {
	var admin model.Admin
	admin.Username = username
	admin.Password = password
	// 判断用户是否存在
	database.Db.Where("username = ?", username).First(&admin)
	// 如果用户存在，修改密码;如果用户不存在，返回false
	if admin.ID != 0 {
		database.Db.Model(&admin).Update("password", password)
		return true
	} else {
		return false
	}
}

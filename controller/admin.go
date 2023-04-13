package controller

import (
	"fmt"

	"example.com/blog/model"
	"example.com/blog/utils"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// Test 测试接口
func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// Login 登录接口
func Login(c *gin.Context) {
	var auth model.Auth
	err := c.BindJSON(&auth)
	if err != nil {
		return
	}
	valid := validation.Validation{}
	ok, _ := valid.Valid(&auth)
	data := make(map[string]interface{})
	if ok {
		isExist := utils.CheckAuth(auth.Username, auth.Password)
		// 如果是登录后管，需要检查这个用户对应的is_admin字段是否为1
		if c.Request.Header.Get("X-from") != "web" {
			if !utils.CheckAdmin(auth.Username) {
				c.JSON(200, gin.H{
					"code": 100,
					"msg":  "该用户不是管理员",
				})
				return
			}
		}
		//	如果在数据库中存在该用户，且密码正确，那么就生成token，否则就返回错误信息
		if isExist {
			token, err := utils.GenerateToken(auth.Username, auth.Password)
			if err != nil {
				c.JSON(200, gin.H{
					"code": 500,
					"msg":  "token生成失败",
					"data": nil,
				})
				return
			}
			data["token"] = token
			data["username"] = auth.Username
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "登录成功",
				"data": data,
			})
		} else {
			//	如果在数据库中不存在该用户，或者密码错误，那么就返回错误信息
			c.JSON(200, gin.H{
				"code": 100,
				"msg":  "用户名或密码错误",
			})
		}
	}
}

// Register 注册接口
func Register(c *gin.Context) {
	var auth model.Auth
	err := c.BindJSON(&auth)
	if err != nil {
		return
	}
	valid := validation.Validation{}
	ok, _ := valid.Valid(&auth)
	if ok {
		isExist := utils.AddUser(auth.Username, auth.Password)
		fmt.Println("isExist", isExist)
		if isExist {
			c.JSON(200, gin.H{
				"code": 409,
				"msg":  "该用户已存在",
				"data": nil,
			})
		} else {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "注册成功",
				"data": auth,
			})
		}
	}
}

// Logout 登出接口
func Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"data": nil,
		"msg":  "退出登录成功",
	})
}

// GetUserInfo 获取github用户信息
func GetUserInfo(c *gin.Context) {
	code := c.Query("code")
	var userInfo = make(map[string]interface{})
	if code != "" {
		// 通过 code, 获取 token
		var tokenAuthUrl, err1 = utils.GetToken(code)
		if err1 != nil {
			fmt.Println("GetTokenErr:", err1)
		}
		// 通过 token, 获取用户信息
		user, err2 := utils.GetUserInfo(tokenAuthUrl)
		if err2 != nil {
			fmt.Println("GetUserInfoErr:", err2)
		}
		userInfo["avatar"] = user["avatar_url"]
		userInfo["username"] = user["login"]
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "获取用户信息成功",
		"data": gin.H{
			"userInfo": userInfo,
		},
	})
}

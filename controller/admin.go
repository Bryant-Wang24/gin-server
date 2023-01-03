package controller

import (
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
		isExist := utils.CheckAuth(auth.UserName, auth.Password)
		//	如果在数据库中存在该用户，且密码正确，那么就生成token，否则就返回错误信息
		if isExist {
			token, err := utils.GenerateToken(auth.UserName, auth.Password)
			if err != nil {
				c.JSON(200, gin.H{
					"code": 500,
					"msg":  "token生成失败",
					"data": nil,
				})
				return
			}
			data["token"] = token
			data["userName"] = auth.UserName
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

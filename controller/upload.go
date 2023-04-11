package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// UploadFile 上传文件
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, gin.H{
			"code": 100,
			"msg":  "上传文件失败",
		})
		return
	}
	// 保存文件
	err = c.SaveUploadedFile(file, "/images/logo/"+file.Filename)
	if err != nil {
		fmt.Println("上传文件失败", err)
		c.JSON(200, gin.H{
			"code": 100,
			"msg":  "上传文件失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "上传文件成功",
		"data": []gin.H{
			{"url": "https://wangqiushuang.online:8080/logo/" + file.Filename},
		},
	})
}

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// UploadFile 上传文件
func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	fmt.Println(file.Filename)
}

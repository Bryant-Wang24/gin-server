package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetIntroduction(c *gin.Context) {
	fmt.Println("获取")
}

func AddIntroduction(c *gin.Context) {
	fmt.Println("添加")
}

func UpdateIntroduction(c *gin.Context) {
	fmt.Println("修改")
}

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// CreateArticle 创建文章
func CreateArticle(c *gin.Context) {
	fmt.Println("创建文章")
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	fmt.Println("删除文章")
}

// UpdateArticle 修改文章
func UpdateArticle(c *gin.Context) {
	fmt.Println("修改文章")
}

// GetArticleList 获取文章列表
func GetArticleList(c *gin.Context) {
	fmt.Println("获取文章列表")
}

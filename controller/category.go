package controller

import (
	"fmt"

	"example.com/blog/database"
	"example.com/blog/model"
	"github.com/gin-gonic/gin"
)

// CreateCategory 创建分类
func CreateCategory(c *gin.Context) {
	var category model.Category
	err := c.BindJSON(&category)
	if err != nil {
		return
	}
	database.Db.Where("name = ?", category.Name).First(&category)
	if category.ID != 0 {
		c.JSON(200, gin.H{
			"code": 100,
			"msg":  "分类名已存在",
		})
		return
	}
	category = model.Category{
		Name: category.Name,
	}
	database.Db.Exec("INSERT INTO categories (name) VALUES (?)", category.Name)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "添加分类成功",
		"data": category,
	})
}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	database.Db.Exec("DELETE FROM categories WHERE id = ?", id)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "删除分类成功",
	})
}

// UpdateCategory 修改分类
func UpdateCategory(c *gin.Context) {
	var category model.Category
	var categoryUpdate model.UpdateCategory
	err := c.BindJSON(&categoryUpdate)
	if err != nil {
		return
	}
	database.Db.Where("name = ?", categoryUpdate.Name).First(&category)
	if category.ID != 0 {
		c.JSON(200, gin.H{
			"code": 100,
			"msg":  "分类名已存在",
		})
		return
	}
	database.Db.Exec("UPDATE categories SET name = ? WHERE id = ?", categoryUpdate.Name, categoryUpdate.ID)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "修改分类成功",
	})
}

// GetCategoryList 获取分类列表
func GetCategoryList(c *gin.Context) {
	//根据name模糊查询
	name := c.Query("name")
	page := c.DefaultQuery("page", "1")          //默认值为1
	pageSize := c.DefaultQuery("pageSize", "20") //默认值为20
	var pageInt, pageSizeInt int
	_, err := fmt.Sscanf(page, "%d", &pageInt)
	if err != nil {
		return
	}
	_, err = fmt.Sscanf(pageSize, "%d", &pageSizeInt)
	if err != nil {
		return
	}
	var category []model.Category
	var totalCount int64
	var article []model.Article
	database.Db.Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&category)
	database.Db.Model(&model.Category{}).Count(&totalCount)
	for i := 0; i < len(category); i++ {
		// 查询分类下的文章数量
		database.Db.Where("categories = ?", category[i].Name).Find(&article)
		category[i].ArticleNum = len(article)
	}
	if name != "" {
		database.Db.Where("name LIKE ?", "%"+name+"%").Find(&category)
		database.Db.Model(&model.Category{}).Where("name LIKE ?", "%"+name+"%").Count(&totalCount)
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "获取分类列表成功",
		"data": gin.H{
			"page":       pageInt,
			"pageSize":   pageSizeInt,
			"totalCount": totalCount,
			"list":       category,
		},
	})
}

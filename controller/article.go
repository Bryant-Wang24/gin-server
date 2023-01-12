package controller

import (
	"example.com/blog/database"
	"example.com/blog/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

var DB *gorm.DB

// CreateArticle 创建文章
func CreateArticle(c *gin.Context) {
	var article model.Article
	err := c.BindJSON(&article)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	article.CreateTime = time.Now()
	article.UpdateTime = time.Now()
	database.Db.Create(&article)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "添加文章成功",
		"data": article,
	})
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
	DB = database.Db
	title := c.Query("title")
	categories := c.Query("categories")
	tags := c.Query("tags")
	status := c.Query("status")
	publishStatus := c.Query("publishStatus")
	createStartTime := c.Query("createStartTime")
	createEndTime := c.Query("createEndTime")
	updateStartTime := c.Query("updateStartTime")
	updateEndTime := c.Query("updateEndTime")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "20")
	//转换为int类型
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)
	var totalCount int64
	var articles []model.Article
	if title != "" {
		DB = DB.Where("title LIKE ?", "%"+title+"%").Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&articles)
		DB = DB.Model(&model.Article{}).Where("title LIKE ?", "%"+title+"%").Count(&totalCount)
	}
	if categories != "" {
		DB = DB.Where("categories LIKE ?", "%"+categories+"%").Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&articles)
		DB = DB.Model(&model.Article{}).Where("categories LIKE ?", "%"+categories+"%").Count(&totalCount)
	}
	if tags != "" {
		fmt.Println("tags", tags)
	}

	if status != "" {
		DB = DB.Where("status = ?", status).Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&articles)
		DB = DB.Model(&model.Article{}).Where("status = ?", status).Count(&totalCount)
	}
	if publishStatus != "" {
		DB = DB.Where("publish_status = ?", publishStatus).Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&articles)
		DB = DB.Model(&model.Article{}).Where("publish_status = ?", publishStatus).Count(&totalCount)
	}
	if createStartTime != "" && createEndTime != "" {
		DB = DB.Where("create_time BETWEEN ? AND ?", createStartTime, createEndTime).Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&articles)
		DB = DB.Model(&model.Article{}).Where("create_time BETWEEN ? AND ?", createStartTime, createEndTime).Count(&totalCount)
	}
	if updateStartTime != "" && updateEndTime != "" {
		DB = DB.Where("update_time BETWEEN ? AND ?", updateStartTime, updateEndTime).Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&articles)
		DB = DB.Model(&model.Article{}).Where("update_time BETWEEN ? AND ?", updateStartTime, updateEndTime).Count(&totalCount)
	}
	//分页查询
	if title == "" && categories == "" && tags == "" && status == "" && publishStatus == "" && createStartTime == "" && createEndTime == "" && updateStartTime == "" && updateEndTime == "" {
		database.Db.Model(&model.Article{}).Count(&totalCount)
		database.Db.Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&articles)
	}
	// 根据article的tags字段，查询出对应的tag name
	for i := 0; i < len(articles); i++ {
		var tagIds []string
		if articles[i].Tags != "" {
			tagIds = strings.Split(articles[i].Tags, ",")
			var tagNames []string
			for j := 0; j < len(tagIds); j++ {
				var tag model.Tag
				database.Db.Where("id = ?", tagIds[j]).First(&tag)
				tagNames = append(tagNames, tag.Name)
			}
			articles[i].Tags = strings.Join(tagNames, ",")
		}
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "获取文章列表成功",
		"data": gin.H{
			"page":       pageInt,
			"pageSize":   pageSizeInt,
			"totalCount": totalCount,
			"list":       articles,
		},
	})
}

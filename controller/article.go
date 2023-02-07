package controller

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"example.com/blog/database"
	"example.com/blog/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	id := c.Param("id")
	database.Db.Exec("DELETE FROM articles WHERE id = ?", id)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "删除文章成功",
	})
}

// UpdateArticle 修改文章
func UpdateArticle(c *gin.Context) {
	var article model.Article
	err := c.BindJSON(&article)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	article.UpdateTime = time.Now()
	database.Db.Model(&article).Where("id = ?", article.ID).Updates(article)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "修改文章成功",
		"data": article,
	})
}

// UpdateArticleStatus 修改文章状态
func UpdateArticleStatus(c *gin.Context) {
	var article model.Article
	err := c.BindJSON(&article)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	database.Db.Model(&article).Where("id = ?", article.ID).Update("status", article.Status)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "修改文章状态成功",
		"data": nil,
	})
}

// UpdateArticlePublishStatus 修改文章发布状态
func UpdateArticlePublishStatus(c *gin.Context) {
	var article model.Article
	err := c.BindJSON(&article)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	database.Db.Model(&article).Where("id = ?", article.ID).Update("publish_status", article.PublishStatus)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "修改文章发布状态成功",
		"data": nil,
	})
}

// 是否开启文章一键收藏
func UpdateArticleCollectStatus(c *gin.Context) {
	type IsCollect struct {
		IsCollect int `json:"isCollect"`
	}
	var isCollect IsCollect
	err := c.BindJSON(&isCollect)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	if isCollect.IsCollect == 1 {
		// 把所有文章的收藏状态改为开启
		database.Db.Exec("UPDATE articles SET is_collect = 1")
	} else {
		// 把所有文章的收藏状态改为关闭
		database.Db.Exec("UPDATE articles SET is_collect = 2")
	}
	var msg string
	if isCollect.IsCollect == 1 {
		msg = "一键开启收藏成功"
	} else {
		msg = "一键关闭收藏成功"
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  msg,
		"data": nil,
	})
}

// GetArticleList 查询单篇文章
func GetSingleArticle(c *gin.Context) {
	id := c.Param("id")
	from := c.Query("type")
	article := model.Article{}
	database.Db.Where("id = ?", id).First(&article)
	// 如果是官网发来的请求，需要把标签id转换成标签名称
	if from == "web" {
		var tagNames []string
		if article.Tags != "" {
			for _, tagId := range strings.Split(article.Tags, ",") {
				var tag model.Tag
				database.Db.Where("id = ?", tagId).First(&tag)
				tagNames = append(tagNames, tag.Name)
			}
			article.Tags = strings.Join(tagNames, ",")
		}
		// 如果是官网发来的请求，需要统计文章的阅读量，每次阅读加1
		// 这里先暂时直接数据库统计，后续可以用redis来实现进行优化
		database.Db.Model(&article).Where("id = ?", article.ID).Update("views", article.Views+1)
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "获取文章成功",
		"data": article,
	})
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
		var tagIds []string
		var tagsArr []string
		tagsArr = strings.Split(tags, ",")
		for i := 0; i < len(tagsArr); i++ {
			var tag model.Tag
			database.Db.Where("name = ?", tagsArr[i]).First(&tag)
			tagIds = append(tagIds, strconv.Itoa(tag.ID))
		}
		DB = DB.Where("tags regexp ?", strings.Join(tagIds, ",")).Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&articles)
		DB = DB.Model(&model.Article{}).Where("tags regexp ?", strings.Join(tagIds, ",")).Count(&totalCount)

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

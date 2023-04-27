package controller

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"example.com/blog/database"
	"example.com/blog/model"
	"example.com/blog/utils"
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
		// 获取请求来源的ip地址
		ip := c.ClientIP()
		ipAndId := ip + id
		// 把拼接好的字符串进行md5加密，用来作为redis的key
		key := utils.MD5(ipAndId)
		// 判断redis中是否存在这个key，如果存在，说明是同一个ip地址访问的，不需要统计阅读量，如果不存在，说明是不同的ip地址访问的，需要统计阅读量
		if !utils.IpExists(key) {
			// 把这个key存入redis中，设置过期时间为1小时，过期时间到了，这个key就会被删除
			database.Rdb.SetNX(context.Background(), key, 1, 1*time.Hour)
			database.Db.Model(&article).Where("id = ?", article.ID).Update("views", article.Views+1)
		}
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
		DB = DB.Model(&model.Article{}).Where("title LIKE ?", "%"+title+"%").Count(&totalCount)
		DB = DB.Where("title LIKE ?", "%"+title+"%").Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&articles)
	}
	if categories != "" {
		DB = DB.Model(&model.Article{}).Where("categories = ?", categories).Count(&totalCount)
		DB = DB.Where("categories = ?", categories).Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&articles)
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
		DB = DB.Model(&model.Article{}).Where("tags regexp ?", strings.Join(tagIds, ",")).Count(&totalCount)
		DB = DB.Where("tags regexp ?", strings.Join(tagIds, ",")).Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&articles)
	}

	if status != "" {
		DB = DB.Model(&model.Article{}).Where("status = ?", status).Count(&totalCount)
		DB = DB.Where("status = ?", status).Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&articles)
	}
	if publishStatus != "" {
		DB = DB.Model(&model.Article{}).Where("publish_status = ?", publishStatus).Count(&totalCount)
		DB = DB.Where("publish_status = ?", publishStatus).Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&articles)
	}
	if createStartTime != "" && createEndTime != "" {
		DB = DB.Model(&model.Article{}).Where("create_time BETWEEN ? AND ?", createStartTime, createEndTime).Count(&totalCount)
		DB = DB.Where("create_time BETWEEN ? AND ?", createStartTime, createEndTime).Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&articles)
	}
	if updateStartTime != "" && updateEndTime != "" {
		DB = DB.Model(&model.Article{}).Where("update_time BETWEEN ? AND ?", updateStartTime, updateEndTime).Count(&totalCount)
		DB = DB.Where("update_time BETWEEN ? AND ?", updateStartTime, updateEndTime).Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&articles)
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

// 点赞文章
func LikeArticle(c *gin.Context) {
	var like model.Like
	err := c.BindJSON(&like)
	if err != nil {
		log.Fatal("like bind json error", err)
	}
	// 如果表里面有这个userId和articleId的记录，说明已经点赞过了
	var likeRecord model.Like
	database.Db.Where("user_id = ? AND article_id = ?", like.UserId, like.ArticleId).First(&likeRecord)
	if likeRecord.ID != 0 {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "已点赞过了",
		})
		return
	}
	// 如果没有点赞过，就插入一条记录
	like.CreateTime = time.Now()
	like.UpdateTime = time.Now()
	database.Db.Create(&like)
	// 更新文章点赞数，查询likes表里面articleId的个数
	var article model.Article
	database.Db.Where("id = ?", like.ArticleId).First(&article)
	var likeCount int64
	database.Db.Model(&model.Like{}).Where("article_id = ?", like.ArticleId).Count(&likeCount)
	article.Like = int(likeCount)
	database.Db.Save(&article)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "点赞成功",
	})
}

// 收藏文章
func CollectArticle(c *gin.Context) {
	var collect model.Collect
	err := c.BindJSON(&collect)
	if err != nil {
		log.Fatal("collect bind json error", err)
	}
	// 如果表里面有这个userId和articleId的记录，说明已经收藏过了
	var collectRecord model.Collect
	database.Db.Where("user_id = ? AND article_id = ?", collect.UserId, collect.ArticleId).First(&collectRecord)
	if collectRecord.ID != 0 {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "已收藏过了",
		})
		return
	}
	// 如果没有收藏过，就插入一条记录
	collect.CreateTime = time.Now()
	collect.UpdateTime = time.Now()
	database.Db.Create(&collect)
	// 更新文章收藏数，查询collects表里面articleId的个数
	var article model.Article
	database.Db.Where("id = ?", collect.ArticleId).First(&article)
	var collectCount int64
	database.Db.Model(&model.Collect{}).Where("article_id = ?", collect.ArticleId).Count(&collectCount)
	article.Collect = int(collectCount)
	database.Db.Save(&article)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "收藏成功",
	})
}

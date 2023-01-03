package controller

import (
	"example.com/blog/database"
	"example.com/blog/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// CreateTag 创建标签
func CreateTag(c *gin.Context) {
	var tag model.Tag
	err := c.BindJSON(&tag)
	if err != nil {
		return
	}
	//如果tag的name在数据库中已经存在，那么就返回错误信息
	database.Db.Where("name = ?", tag.Name).First(&tag)
	if tag.ID != 0 {
		c.JSON(200, gin.H{
			"code": 100,
			"msg":  "标签名已存在",
		})
		return
	}
	//如果tag的name在数据库中不存在，那么就将tag插入到数据库中
	tag = model.Tag{
		Name:       tag.Name,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	database.Db.Create(&tag)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "添加标签成功",
		"data": tag,
	})
}

// DeleteTag 删除标签
func DeleteTag(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete tag",
	})
}

// UpdateTag 更新标签
func UpdateTag(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "update tag",
	})
}

// GetTagList 获取标签列表
func GetTagList(c *gin.Context) {
	//根据name模糊查询
	name := c.Query("name")
	page := c.DefaultQuery("page", "1")          //默认值为1
	pageSize := c.DefaultQuery("pageSize", "20") //默认值为20
	fmt.Println(name, page, pageSize)
	//把page和pageSize转换为int类型
	var pageInt, pageSizeInt int
	_, err := fmt.Sscanf(page, "%d", &pageInt)
	if err != nil {
		return
	}
	_, err = fmt.Sscanf(pageSize, "%d", &pageSizeInt)
	if err != nil {
		return
	}
	var tags []model.Tag
	var totalCount int64
	//分页查询数据库，按照创建时间排序,翻译为sql语句为：SELECT * FROM `tags`  ORDER BY `tags`.`created_at` DESC LIMIT 10 OFFSET 0
	database.Db.Order("create_time desc").Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&tags)
	// 查询数据库总数 翻译为sql语句为：SELECT count(*) FROM `tags`
	database.Db.Model(&model.Tag{}).Count(&totalCount)
	//如果传过来name不为空，就模糊查询
	if name != "" {
		database.Db.Where("name LIKE ?", "%"+name+"%").Find(&tags)
		// 查询数据库总数 翻译为sql语句为：SELECT count(*) FROM `tags`  WHERE (name LIKE '%a%')
		database.Db.Model(&model.Tag{}).Where("name LIKE ?", "%"+name+"%").Count(&totalCount)
	}
	//返回的数据放在data中
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "获取标签列表成功",
		"data": gin.H{
			"page":       pageInt,
			"pageSize":   pageSizeInt,
			"totalCount": totalCount,
			"list":       tags,
		},
	})
}

// UpdateTagStatus 更新标签状态
func UpdateTagStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "update tag status",
	})
}

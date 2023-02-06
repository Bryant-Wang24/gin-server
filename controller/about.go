package controller

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"example.com/blog/database"
	"example.com/blog/model"
	"github.com/gin-gonic/gin"
)

// GetAbout 获取关于信息
func GetAbout(c *gin.Context) {
	var about model.About
	// 查询所有数据
	database.Db.Find(&about)
	// 如果没有数据，返回空
	if about.ID == 0 {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "获取关于信息成功",
			"data": nil,
		})
		return
	}

	// 把Tags转换为数组
	var tagsArr []string
	tagsArr = strings.Split(about.Tags, ",")
	var imageIds []int
	// 把图片id转换成int类型
	for _, id := range strings.Split(about.Imgs, ",") {
		imageId, _ := strconv.Atoi(id)
		imageIds = append(imageIds, imageId)
	}
	// 从images表中根据id分别查询出图片的信息
	var images []model.Image
	for _, id := range imageIds {
		var image model.Image
		// 根据id查询图片信息
		database.Db.Where("id = ?", id).Find(&image)
		images = append(images, image)
	}
	aboutInfo := model.AboutInfo{
		ID:         about.ID,
		Tags:       tagsArr,
		Desc:       about.Desc,
		ShowResume: about.ShowResume,
		CreateTime: about.CreateTime,
		UpdateTime: about.UpdateTime,
		Imgs:       images,
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "获取关于信息成功",
		"data": aboutInfo,
	})
}

// 添加关于信息
func AddAbout(c *gin.Context) {
	var about model.About
	var aboutInfo model.AboutInfo

	err := c.BindJSON(&aboutInfo)
	if err != nil {
		fmt.Println("添加关于信息失败", err)
		return
	}
	for _, image := range aboutInfo.Imgs {
		var images model.Image
		images = image
		// 向images表中添加图片信息
		database.Db.Create(&images)
	}
	// 向about表中添加关于信息
	about.Tags = strings.Join(aboutInfo.Tags, ",")
	about.Desc = aboutInfo.Desc
	about.ShowResume = aboutInfo.ShowResume
	about.CreateTime = time.Now()
	about.UpdateTime = time.Now()
	// 向about表imgs字段中添加图片id,id从images表中获取
	// 查询images表中所有的ID
	var images []model.Image
	database.Db.Find(&images)
	var imageIds []string
	for _, image := range images {
		imageIds = append(imageIds, strconv.Itoa(image.ID))
	}
	about.Imgs = strings.Join(imageIds, ",")
	database.Db.Create(&about)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "添加关于信息成功",
		"data": aboutInfo,
	})
}

// 修改关于信息
func UpdateAbout(c *gin.Context) {
	var about model.About
	var aboutInfo model.AboutInfo
	err := c.BindJSON(&aboutInfo)
	if err != nil {
		return
	}
	//硬删除images表中所有的数据
	database.Db.Unscoped().Where("id > ?", 0).Delete(&model.Image{})
	// 添加新的图片信息
	for _, image := range aboutInfo.Imgs {
		var images model.Image
		images = image
		// 向images表中添加图片信息
		database.Db.Create(&images)
	}
	// 更新about表中对应的id的关于信息
	database.Db.Model(&about).Where("id = ?", aboutInfo.ID).Updates(model.About{
		Tags:       strings.Join(aboutInfo.Tags, ","),
		Desc:       aboutInfo.Desc,
		ShowResume: aboutInfo.ShowResume,
		UpdateTime: time.Now(),
	})
	// 查询images表中所有的ID
	var images []model.Image
	database.Db.Find(&images)
	var imageIds []string
	for _, image := range images {
		imageIds = append(imageIds, strconv.Itoa(image.ID))
	}
	// 更新about表中对应的id的关于信息
	database.Db.Model(&about).Where("id = ?", aboutInfo.ID).Updates(model.About{
		Imgs: strings.Join(imageIds, ","),
	})
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "修改关于信息成功",
		"data": aboutInfo,
	})
}

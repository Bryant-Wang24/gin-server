package controller

import (
	"strconv"
	"strings"

	"example.com/blog/database"
	"example.com/blog/model"
	"github.com/gin-gonic/gin"
)

// 获取关于信息
func GetAbout(c *gin.Context) {
	var about model.About
	// 查询所有数据
	database.Db.Find(&about)
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

// 修改关于信息
func UpdateAbout(c *gin.Context) {
	var about model.About
	var aboutInfo model.AboutInfo
	err := c.BindJSON(&aboutInfo)
	if err != nil {
		return
	}
	// 把Tags转换为字符串
	about.Tags = strings.Join(aboutInfo.Tags, ",")
	// 把图片id转换成字符串
	var imageIds []string
	for _, image := range aboutInfo.Imgs {
		imageIds = append(imageIds, strconv.Itoa(image.ID))
	}
	about.Imgs = strings.Join(imageIds, ",")
	// 更新images表中对应的id的图片信息
	for _, image := range aboutInfo.Imgs {
		// 创建一个images的结构体
		var images model.Image
		// 把aboutInfo中的图片信息赋值给images
		images = image
		// 更新images表中对应的id的图片信息
		database.Db.Model(&images).Where("id = ?", image.ID).Updates(image)
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "修改关于信息成功",
		"data": aboutInfo,
	})
}

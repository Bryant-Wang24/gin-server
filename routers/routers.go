package routers

import (
	"example.com/blog/controller"
	"example.com/blog/middleware/jwt"
	"github.com/gin-gonic/gin"
)

var V1 = "api/v1"

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET(V1+"/ping", controller.Test)
	r.POST(V1+"/admin/login", controller.Login)
	r.POST(V1+"/admin/logout", controller.Logout)
	//标签的路由组
	tag := r.Group(V1 + "/tags").Use(jwt.JWT())
	{
		tag.POST("", controller.CreateTag)
		tag.DELETE("/:id", controller.DeleteTag)
		tag.PUT("", controller.UpdateTag)
		tag.GET("", controller.GetTagList)
		//启用停用标签
		tag.PUT("/status", controller.UpdateTagStatus)
	}
	//分类的路由组
	category := r.Group(V1 + "/categories").Use(jwt.JWT())
	{
		category.POST("", controller.CreateCategory)
		category.DELETE("/:id", controller.DeleteCategory)
		category.PUT("", controller.UpdateCategory)
		category.GET("", controller.GetCategoryList)
	}

	//上传文件
	r.POST(V1+"/upload", controller.UploadFile)

	err := r.Run(":7001")
	if err != nil {
		return nil
	} // listen and serve on
	return r
}

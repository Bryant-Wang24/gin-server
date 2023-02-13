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
	r.POST(V1+"/admin/register", controller.Register)
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

	//文章的路由组
	article := r.Group(V1 + "/articles").Use(jwt.JWT())
	{
		article.POST("", controller.CreateArticle)
		article.DELETE("/:id", controller.DeleteArticle)
		article.GET("", controller.GetArticleList)
		// 查询单篇文章
		article.GET("/:id/edit", controller.GetSingleArticle)
		// 修改文章
		article.PUT("", controller.UpdateArticle)
		// 修改文章状态
		article.PUT("/status", controller.UpdateArticleStatus)
		// 修改文章发布状态
		article.PUT("/publishStatus", controller.UpdateArticlePublishStatus)
		// 是否开启文章一键收藏
		article.POST("/collectStatus", controller.UpdateArticleCollectStatus)
	}

	//个人简介路由组
	introduction := r.Group(V1 + "/config/right/introduction").Use(jwt.JWT())
	{
		introduction.GET("", controller.GetIntroduction)
		introduction.POST("", controller.AddIntroduction)
		introduction.PUT("", controller.UpdateIntroduction)
	}

	//上传文件
	r.POST(V1+"/upload", controller.UploadFile)

	// 关于路由组
	about := r.Group(V1 + "/about").Use(jwt.JWT())
	{
		about.GET("", controller.GetAbout)
		about.PUT("", controller.UpdateAbout)
		about.POST("", controller.AddAbout)
	}

	err := r.Run(":7001")
	if err != nil {
		return nil
	} // listen and serve on
	return r
}

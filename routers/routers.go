package routers

import (
	"example.com/blog/controller"
	"github.com/gin-gonic/gin"
)

var V1 = "api/v1"

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET(V1+"/ping", controller.Test)
	r.POST(V1+"/admin/login", controller.Login)
	r.POST(V1+"/admin/logout", controller.Logout)
	err := r.Run(":7001")
	if err != nil {
		return nil
	} // listen and serve on
	return r
}

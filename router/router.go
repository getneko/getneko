package router

import (
	"getneko/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/chacklowversion", controller.ChackLowVersion)
		v1.POST("/userreg", controller.Uerregcontroller)
		v1.POST("/userlogin", controller.UerLogincontroller)
		v1.POST("/createproject", controller.Projectcreatecontroller)
	}

	return r
}

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
		v1.POST("/delproject", controller.Projectdelcontroller)
		v1.GET("/usersearch", controller.Uernamesearchcontroller)
		v1.POST("/getpermissionlist", controller.Getpermissionlistcontroller)
		v1.POST("/setpermission", controller.Setpermissioncontroller)
		v1.POST("/addapi", controller.Addapicontroller)
		v1.POST("/getprojectlist", controller.GetProjectlistcontroller)
	}

	return r
}

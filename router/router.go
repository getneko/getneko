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
	}

	return r
}

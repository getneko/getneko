package controller

import (
	"getneko/structtypes"

	"github.com/gin-gonic/gin"
)

// @Summary 获取客户端最低版本
// @Description 比较客户端与服务器的版本号
// @Tags 客户端信息
// @Accept json
// @Produce json
// @Success 200 {object} structtypes.JSONResult{data=string} "desc"
// @Router /v1/chacklowversion [get]
func ChackLowVersion(c *gin.Context) {
	c.JSON(200, &structtypes.JSONResult{Code: 0, Message: "", Data: "1.0"})
}

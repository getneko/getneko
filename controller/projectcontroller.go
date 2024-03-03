package controller

import (
	"getneko/db"
	"getneko/structtypes"
	"getneko/tool"

	"github.com/gin-gonic/gin"
)

// @Summary 创建项目
// @Description 目前暂时没有做权限
// @Tags 项目操作
// @Accept json
// @Produce json
// @Param UserLogin  body structtypes.Createproject true "project info"
// @Success 200 {object} structtypes.JSONResult{data=string} "desc"
// @Router /v1/createproject [post]
func Projectcreatecontroller(c *gin.Context) {
	//参数校验
	var procreatereq structtypes.Createproject
	if err := c.ShouldBindJSON(&procreatereq); err != nil {
		c.JSON(200, tool.Refal(-1, "illegal request"))
		return
	}
	//检测用户登录状态
	users := tool.Chackuserlogin(procreatereq.Username, procreatereq.Tokens)
	if users == nil {
		c.JSON(200, tool.Refal(-4, "user not login"))
		return
	}
	//创建项目
	db.ORMDB.Create(&structtypes.Project{Name: procreatereq.Name, Createuser: procreatereq.Username})
	c.JSON(200, tool.ReSucc("Created successfully"))
}

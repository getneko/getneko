package controller

import (
	"getneko/db"
	"getneko/structtypes"
	"getneko/tool"
	"strconv"

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
	//检测项目是否存在
	var projects structtypes.Project
	err := db.ORMDB.Where("createuser = ? and name = ?", users.ID, procreatereq.Name).First(&projects).Error
	if err == nil {
		c.JSON(200, tool.Refal(-6, "project is exist"))
		return
	}
	//创建项目
	db.ORMDB.Create(&structtypes.Project{Name: procreatereq.Name, Createuser: strconv.Itoa(int(users.ID))})
	//添加管理员权限
	var crtpro structtypes.Project
	db.ORMDB.Where("createuser = ? and name = ?", users.ID, procreatereq.Name).First(&crtpro)
	db.ORMDB.Create(&structtypes.Permissions{Userid: int(users.ID), Levels: 3, Projectid: crtpro.ID})
	c.JSON(200, tool.ReSucc("Created successfully"))
}

// @Summary 删除项目
// @Description 只要创建者才能删除
// @Tags 项目操作
// @Accept json
// @Produce json
// @Param UserLogin  body structtypes.Delproject true "project info"
// @Success 200 {object} structtypes.JSONResult{data=string} "desc"
// @Router /v1/delproject [post]
func Projectdelcontroller(c *gin.Context) {
	//参数校验
	var projectdelreq structtypes.Delproject
	if err := c.ShouldBindJSON(&projectdelreq); err != nil {
		c.JSON(200, tool.Refal(-1, "illegal request"))
		return
	}
	//检测用户登录状态
	users := tool.Chackuserlogin(projectdelreq.Username, projectdelreq.Tokens)
	if users == nil {
		c.JSON(200, tool.Refal(-4, "user not login"))
		return
	}
	//检测项目是否存在
	var projects structtypes.Project
	err := db.ORMDB.Where("createuser = ? and name = ?", users.ID, projectdelreq.Name).First(&projects).Error
	if err != nil {
		c.JSON(200, tool.Refal(-5, "project does not exist or you are not the creator "))
		return
	}
	//执行删除
	db.ORMDB.Unscoped().Delete(&projects)
	db.ORMDB.Where("projectid = ?", projects.ID).Unscoped().Delete(&structtypes.Permissions{})
	c.JSON(200, tool.ReSucc("Delete successfully"))
}

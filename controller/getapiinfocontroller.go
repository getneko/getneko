package controller

import (
	"getneko/db"
	"getneko/structtypes"
	"getneko/tool"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// @Summary 获取api列表
// @Description 目前没有需要注意的
// @Tags api记录操作
// @Accept json
// @Produce json
// @Param UserLogin  body structtypes.Getapilist true "api info"
// @Success 200 {object} structtypes.JSONResult{data=[]structtypes.Getapilistres} "desc"
// @Router /v1/getapilist [post]
func Getapicontroller(c *gin.Context) {
	//参数校验
	var getapilist structtypes.Getapilist
	if err := c.ShouldBindJSON(&getapilist); err != nil {
		log.Error(err)
		c.JSON(200, tool.Refal(-1, "illegal request"))
		return
	}
	//检测用户登录状态
	users := tool.Chackuserlogin(getapilist.Username, getapilist.Tokens)
	if users == nil {
		c.JSON(200, tool.Refal(-4, "user not login"))
		return
	}
	//检测项目是否存在
	var projects structtypes.Project
	err := db.ORMDB.Where("name = ?", getapilist.Projectname).First(&projects).Error
	if err != nil {
		c.JSON(200, tool.Refal(-7, "project does not exist"))
		return
	}
	//检查是否拥有权限
	if !tool.Chackuserhavepermission(int(users.ID), projects.ID, 0) {
		c.JSON(200, tool.Refal(-8, "User does not have permission"))
		return
	}
	var res []structtypes.Getapilistres
	db.ORMDB.Model(&structtypes.Apirecode{}).Where("projectid = ?", projects.ID).Find(&res)
	c.JSON(200, tool.ReSucc(res))
}

// @Summary 通过id获取api
// @Description 目前没有需要注意的
// @Tags api记录操作
// @Accept json
// @Produce json
// @Param UserLogin  body structtypes.Getapi true "api info"
// @Success 200 {object} structtypes.JSONResult{data=structtypes.Getapires} "desc"
// @Router /v1/getapibyid [post]
func Getapibyidcontroller(c *gin.Context) {
	//参数校验
	var getapi structtypes.Getapi
	if err := c.ShouldBindJSON(&getapi); err != nil {
		log.Error(err)
		c.JSON(200, tool.Refal(-1, "illegal request"))
		return
	}
	//检测用户登录状态
	users := tool.Chackuserlogin(getapi.Username, getapi.Tokens)
	if users == nil {
		c.JSON(200, tool.Refal(-4, "user not login"))
		return
	}
	//检测项目是否存在
	var projects structtypes.Project
	err := db.ORMDB.Where("name = ?", getapi.Projectname).First(&projects).Error
	if err != nil {
		c.JSON(200, tool.Refal(-7, "project does not exist"))
		return
	}
	//检查是否拥有权限
	if !tool.Chackuserhavepermission(int(users.ID), projects.ID, 0) {
		c.JSON(200, tool.Refal(-8, "User does not have permission"))
		return
	}
	var apis []structtypes.Getapires
	db.ORMDB.Model(&structtypes.Apirecode{}).Where("id = ?", getapi.Apiid).Find(&apis)
	if len(apis) == 0 {
		c.JSON(200, tool.Refal(-13, "api does not exist"))
		return
	}
	c.JSON(200, tool.ReSucc(apis[0]))
}

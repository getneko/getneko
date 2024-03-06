package controller

import (
	"getneko/db"
	"getneko/structtypes"
	"getneko/tool"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// @Summary 创建api
// @Description 参数校验慢慢完善
// @Tags api记录操作
// @Accept json
// @Produce json
// @Param UserLogin  body structtypes.Addapireq true "api info"
// @Success 200 {object} structtypes.JSONResult{data=string} "desc"
// @Router /v1/addapi [post]
func Addapicontroller(c *gin.Context) {
	//参数校验
	var addapireq structtypes.Addapireq
	if err := c.ShouldBindJSON(&addapireq); err != nil {
		log.Error(err)
		c.JSON(200, tool.Refal(-1, "illegal request"))
		return
	}
	//检测用户登录状态
	users := tool.Chackuserlogin(addapireq.Username, addapireq.Tokens)
	if users == nil {
		c.JSON(200, tool.Refal(-4, "user not login"))
		return
	}
	//检测项目是否存在
	var projects structtypes.Project
	err := db.ORMDB.Where("name = ?", addapireq.Projectname).First(&projects).Error
	if err != nil {
		c.JSON(200, tool.Refal(-7, "project does not exist"))
		return
	}
	//检查是否拥有权限
	if !tool.Chackuserhavepermission(int(users.ID), projects.ID, 1) {
		c.JSON(200, tool.Refal(-8, "User does not have permission"))
		return
	}
	//写入数据库
	newapi := structtypes.Apirecode{
		Types:        addapireq.Types,
		Path:         addapireq.Path,
		Words:        addapireq.Words,
		Headcode:     addapireq.Headcode,
		Pathcode:     addapireq.Pathcode,
		Bodytype:     addapireq.Bodytype,
		Bodyrecodes:  addapireq.Bodyrecodes,
		Returntype:   addapireq.Returntype,
		Returncodes:  addapireq.Returntype,
		IsDeprecated: addapireq.IsDeprecated,
		Projectid:    projects.ID,
	}
	db.ORMDB.Create(&newapi)
	c.JSON(200, tool.ReSucc("Created successfully"))
}

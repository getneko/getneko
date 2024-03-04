package controller

import (
	"getneko/db"
	"getneko/structtypes"
	"getneko/tool"

	"github.com/gin-gonic/gin"
)

// @Summary 获取项目权限表
// @Description 需要管理员权限
// @Tags 权限操作
// @Accept json
// @Produce json
// @Param UserLogin  body structtypes.Getpermissionlist true "project info"
// @Success 200 {object} structtypes.JSONResult{data=[]structtypes.Getpermissionlistres} "desc"
// @Router /v1/Getpermissionlist [post]
func Getpermissionlistcontroller(c *gin.Context) {
	//参数校验
	var getpermissionlist structtypes.Getpermissionlist
	if err := c.ShouldBindJSON(&getpermissionlist); err != nil {
		c.JSON(200, tool.Refal(-1, "illegal request"))
		return
	}
	//检测用户登录状态
	users := tool.Chackuserlogin(getpermissionlist.Username, getpermissionlist.Tokens)
	if users == nil {
		c.JSON(200, tool.Refal(-4, "user not login"))
		return
	}
	//检测项目是否存在
	var projects structtypes.Project
	err := db.ORMDB.Where("name = ?", getpermissionlist.Projectname).First(&projects).Error
	if err != nil {
		c.JSON(200, tool.Refal(-7, "project does not exist"))
		return
	}
	//获取权限组
	var prm []structtypes.Permissions
	db.ORMDB.Where("projectid = ?", projects.ID).Find(&prm)
	//处理返回值
	var res []structtypes.Getpermissionlistres
	prmflags := true //校验是否有权限
	for i := 0; i < len(prm); i++ {
		if prm[i].Userid == int(users.ID) && prm[i].Levels == 3 {
			prmflags = false
		}
		res = append(res, structtypes.Getpermissionlistres{UserID: prm[i].Userid, Username: tool.Getusernamebyid(prm[i].Userid), Levels: prm[i].Levels})
	}
	if prmflags {
		c.JSON(200, tool.Refal(-8, "User does not have permission"))
	} else {
		c.JSON(200, tool.ReSucc(res))
	}
}

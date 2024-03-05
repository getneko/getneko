package controller

import (
	"getneko/db"
	"getneko/structtypes"
	"getneko/tool"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// @Summary 获取项目权限表
// @Description 需要管理员权限
// @Tags 权限操作
// @Accept json
// @Produce json
// @Param UserLogin  body structtypes.Getpermissionlist true "project info"
// @Success 200 {object} structtypes.JSONResult{data=[]structtypes.Getpermissionlistres} "desc"
// @Router /v1/getpermissionlist [post]
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
		if prm[i].Userid == int(users.ID) && prm[i].Levels == 2 {
			prmflags = false
		}
		res = append(res, structtypes.Getpermissionlistres{UserID: prm[i].Userid, Username: tool.Getusernamebyid(prm[i].Userid), Levels: prm[i].Levels})
	}
	if prmflags {
		c.JSON(200, tool.Refal(-8, "User does not have permission"))
		return
	} else {
		c.JSON(200, tool.ReSucc(res))
	}
}

// @Summary 设置权限
// @Description 需要管理员权限
// @Tags 权限操作
// @Accept json
// @Produce json
// @Param UserLogin  body structtypes.Setpermission true "project info"
// @Success 200 {object} structtypes.JSONResult{data=string} "desc"
// @Router /v1/setpermission [post]
func Setpermissioncontroller(c *gin.Context) {
	//参数校验
	var setpermission structtypes.Setpermission
	if err := c.ShouldBindJSON(&setpermission); err != nil {
		log.Error(err)
		c.JSON(200, tool.Refal(-1, "illegal request"))
		return
	}
	//检测用户登录状态
	users := tool.Chackuserlogin(setpermission.Username, setpermission.Tokens)
	if users == nil {
		c.JSON(200, tool.Refal(-4, "user not login"))
		return
	}
	//检测项目是否存在
	var projects structtypes.Project
	err := db.ORMDB.Where("name = ?", setpermission.Projectname).First(&projects).Error
	if err != nil {
		c.JSON(200, tool.Refal(-7, "project does not exist"))
		return
	}
	//检查是否为管理员
	var chackpri structtypes.Permissions
	err = db.ORMDB.Where("userid = ? and projectid = ? and levels =2", users.ID, projects.ID).First(&chackpri).Error
	if err != nil {
		c.JSON(200, tool.Refal(-8, "User does not have permission"))
		return
	}
	//清空权限
	db.ORMDB.Where("projectid = ?", projects.ID).Unscoped().Delete(&structtypes.Permissions{})
	//写入权限
	//管理员
	adminarr := strings.Split(setpermission.Adminnames, ",")
	if adminarr[0] != "" {
		var adminid []int
		for i := 0; i < len(adminarr); i++ {
			adminidnow := tool.Getidbyusername(adminarr[i])
			if adminidnow == -1 {
				c.JSON(200, tool.Refal(-9, "user does not exist"))
				return
			} else {
				adminid = append(adminid, adminidnow)
			}
		}
		for i := 0; i < len(adminarr); i++ {
			db.ORMDB.Create(&structtypes.Permissions{Userid: adminid[i], Projectid: projects.ID, Levels: 2})
		}
	}
	//编辑者
	editarr := strings.Split(setpermission.Editnames, ",")
	if editarr[0] != "" {
		var editid []int
		for i := 0; i < len(editarr); i++ {
			editidnow := tool.Getidbyusername(editarr[i])
			if editidnow == -1 {
				c.JSON(200, tool.Refal(-9, "user does not exist"))
				return
			} else {
				editid = append(editid, editidnow)
			}
		}
		for i := 0; i < len(editarr); i++ {
			db.ORMDB.Create(&structtypes.Permissions{Userid: editid[i], Projectid: projects.ID, Levels: 1})
		}
	}
	//访客
	guestarr := strings.Split(setpermission.Guestnames, ",")
	if guestarr[0] != "" {
		var guestid []int
		for i := 0; i < len(guestarr); i++ {
			guestidnow := tool.Getidbyusername(guestarr[i])
			if guestidnow == -1 {
				c.JSON(200, tool.Refal(-9, "user does not exist"))
				return
			} else {
				guestid = append(guestid, guestidnow)
			}
		}
		for i := 0; i < len(guestarr); i++ {
			db.ORMDB.Create(&structtypes.Permissions{Userid: guestid[i], Projectid: projects.ID, Levels: 0})
		}
	}
	c.JSON(200, tool.ReSucc("Setup successful"))
}

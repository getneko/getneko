package controller

import (
	"getneko/db"
	"getneko/structtypes"
	"getneko/tool"

	"github.com/gin-gonic/gin"
)

// @Summary 用户注册
// @Description 用户注册api
// @Tags 用户操作
// @Accept json
// @Produce json
// @Param UserRegReq  body structtypes.UserRegReq true "user info"
// @Success 200 {object} structtypes.JSONResult{data=string} "desc"
// @Router /v1/userreg [post]
func Uerregcontroller(c *gin.Context) {
	//参数校验
	var userregreq structtypes.UserRegReq
	if err := c.ShouldBindJSON(&userregreq); err != nil {
		c.JSON(400, tool.Refal("illegal request"))
		return
	}
	//检查数据是否已经存在
	var users structtypes.User
	err := db.ORMDB.Where("Username = ? or Email = ?", userregreq.Username, userregreq.Email).First(&users).Error
	if err == nil {
		c.JSON(400, tool.Refal("Username or password already exists"))
		return
	}
	//生成盐
	salts := tool.GetRandString(10)
	//生成令牌
	tokens := tool.GetRandString(30)
	//加密密码
	pass := tool.GetSha256(salts + userregreq.Password)
	//写入数据库
	db.ORMDB.Create(&structtypes.User{Username: userregreq.Username, Password: pass, Salt: salts, Language: userregreq.Language, Email: userregreq.Email, Token: tokens})
	c.JSON(200, tool.ReSucc(tokens))
}

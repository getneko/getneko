package tool

import (
	"getneko/db"
	"getneko/structtypes"
)

// 校验用户是否登录
func Chackuserlogin(username, tokens string) *structtypes.User {
	var userdate structtypes.User
	err := db.ORMDB.Where("username = ? and token = ?", username, tokens).First(&userdate).Error
	if err != nil {
		return nil
	} else {
		return &userdate
	}
}

// 通过id查询用户名
func Getusernamebyid(id int) string {
	var userdate structtypes.User
	db.ORMDB.Where("id = ?", id).First(&userdate)
	return userdate.Username
}

// 通过用户名查询id
func Getidbyusername(username string) int {
	var userdata structtypes.User
	err := db.ORMDB.Where("username = ?", username).First(&userdata).Error
	if err != nil {
		return -1
	} else {
		return int(userdata.ID)
	}
}

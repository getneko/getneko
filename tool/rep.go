package tool

import (
	"getneko/db"
	"getneko/structtypes"
)

// 返回错误结果
func Refal(code int, msg string) *structtypes.JSONResult {
	return &structtypes.JSONResult{Code: code, Message: msg, Data: nil}
}

// 返回正常结果
func ReSucc(data interface{}) *structtypes.JSONResult {
	return &structtypes.JSONResult{Code: 0, Message: "", Data: data}
}

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

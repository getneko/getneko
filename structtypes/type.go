package structtypes

import "gorm.io/gorm"

// 通用结构体
type JSONResult struct {
	Code    int         `json:"code" `
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 用户表
type User struct {
	gorm.Model
	Username string //用户名
	Password string //密码
	Salt     string //盐
	Language string //语言
	Email    string //邮箱
	Token    string //令牌
}

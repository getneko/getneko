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

// 项目
type Project struct {
	ID         int    //id
	Name       string //名字
	Createuser string //创建者
	Edituser   string //编辑者
	Guest      string //只读访问者
	Feats      string //保留字段
}

// api 记录
type Apirecode struct {
	ID           int    //id
	Types        string //请求类型
	Path         string //路径
	Words        string //注释
	Headcode     string //请求头部结构
	Pathcode     string //path结构
	Bodytype     string //body的类型
	Bodyrecodes  string //body结构
	Returntype   string //返回类型
	Returncodes  string //返回结构
	IsDeprecated int    //弃用状态
	Feats        string //保留字段
}

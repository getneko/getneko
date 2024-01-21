package tool

import "getneko/structtypes"

// 返回错误结果
func Refal(code int, msg string) *structtypes.JSONResult {
	return &structtypes.JSONResult{Code: code, Message: msg, Data: nil}
}

// 返回正常结果
func ReSucc(data interface{}) *structtypes.JSONResult {
	return &structtypes.JSONResult{Code: 0, Message: "", Data: data}
}

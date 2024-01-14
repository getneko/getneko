package structtypes

// 通用结构体
type JSONResult struct {
	Code    int         `json:"code" `
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

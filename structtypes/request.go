package structtypes

// 用户注册请求结构体
type UserRegReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Language string `json:"language" binding:"required,bcp47_language_tag"`
	Email    string `json:"email" binding:"required,email"`
}

// 用户登陆结构体
type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Language string `json:"language" binding:"required,bcp47_language_tag"`
}
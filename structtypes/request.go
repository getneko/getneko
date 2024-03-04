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

//创建项目结构体
type Createproject struct {
	Name     string `json:"name" binding:"required"`     //名字
	Username string `json:"username" binding:"required"` //用户名
	Tokens   string `json:"tokens" binding:"required"`   //用户token
}

//删除项目结构体
type Delproject struct {
	Name     string `json:"name" binding:"required"`     //名字
	Username string `json:"username" binding:"required"` //用户名
	Tokens   string `json:"tokens" binding:"required"`   //用户token
}

//模糊查询结构体
type Usersearch struct {
	Name string `form:"name" binding:"required"`
}

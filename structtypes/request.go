package structtypes

// 用户注册请求结构体
type UserRegReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Language string `json:"language" binding:"required,oneof=zh-Hans zh-Hant en-US"`
	Email    string `json:"email" binding:"required,email"`
}

// 用户登陆结构体
type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Language string `json:"language" binding:"required,oneof=zh-Hans zh-Hant en-US"`
}

// 创建项目结构体
type Createproject struct {
	Projectname string `json:"projectname" binding:"required"` //名字
	Username    string `json:"username" binding:"required"`    //用户名
	Tokens      string `json:"tokens" binding:"required"`      //用户token
}

// 删除项目结构体
type Delproject struct {
	Projectname string `json:"projectname" binding:"required"` //名字
	Username    string `json:"username" binding:"required"`    //用户名
	Tokens      string `json:"tokens" binding:"required"`      //用户token
}

// 模糊查询结构体
type Usersearch struct {
	Username string `form:"username"`
}

// 查询项目权限列表结构体
type Getpermissionlist struct {
	Projectname string `json:"projectname" binding:"required"` //项目名
	Username    string `json:"username" binding:"required"`    //用户名
	Tokens      string `json:"tokens" binding:"required"`      //用户token
}

// 设置权限结构体
type Setpermission struct {
	Adminnames  string `json:"adminnames"`                     //管理员
	Editnames   string `json:"editnames"`                      //编辑者
	Guestnames  string `json:"guestnames"`                     //访客
	Projectname string `json:"projectname" binding:"required"` //项目名
	Username    string `json:"username" binding:"required"`    //用户名
	Tokens      string `json:"tokens" binding:"required"`      //用户token
}

// 获取项目列表结构体
type GetProjectlist struct {
	Username string `json:"username" binding:"required"` //用户名
	Tokens   string `json:"tokens" binding:"required"`   //用户token
}

// 获取api列表
type Getapilist struct {
	Username    string `json:"username" binding:"required"`    //用户名
	Tokens      string `json:"tokens" binding:"required"`      //用户token
	Projectname string `json:"projectname" binding:"required"` //项目名
}

// 通过id获取api详情
type Getapi struct {
	Username    string `json:"username" binding:"required"`    //用户名
	Tokens      string `json:"tokens" binding:"required"`      //用户token
	Projectname string `json:"projectname" binding:"required"` //项目名
	Apiid       int    `json:"apiid" binding:"required"`
}

// 删除api请求
type Delapi struct {
	Username    string `json:"username" binding:"required"`    //用户名
	Tokens      string `json:"tokens" binding:"required"`      //用户token
	Projectname string `json:"projectname" binding:"required"` //项目名
	Apiid       int    `json:"apiid" binding:"required"`
}

package structtypes

//模糊查询返回
type Usersearchres struct {
	ID       int
	Username string
}

//查询项目权限返回
type Getpermissionlistres struct {
	UserID   int
	Username string
	Levels   int
}

//获取项目列表返回
type Getprojectlistres struct {
	Projectid   int
	Projectname string
	Creater     string
	Apinum      int
}

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

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

//获取api列表返回
type Getapilistres struct {
	ID           int    //id
	Types        string //请求类型
	Path         string //路径
	IsDeprecated int    //弃用状态 1表示正常 2表示弃用
}

//通过if寻找记录返回
type Getapires struct {
	Types        string //请求类型
	Path         string //路径
	Words        string //注释
	Headcode     string //请求头部结构
	Pathcode     string //path结构
	Bodytype     string //body的类型
	Bodyrecodes  string //body结构
	Returntype   string //返回类型
	Returncodes  string //返回结构
	IsDeprecated int    //弃用状态 1表示正常 2表示弃用
}

package structtypes

//因为比较复杂所以我把这部分独立出来了
type Addapireq struct {
	Username     string `json:"username" binding:"required"`               //用户名
	Tokens       string `json:"tokens" binding:"required"`                 //用户token
	Types        string `json:"types" binding:"required,oneof=get post"`   //请求类型
	Path         string `json:"path" binding:"required"`                   //路径
	Words        string `json:"words" binding:"required"`                  //注释
	Projectname  string `json:"projectname" binding:"required"`            //项目名
	Headcode     string `json:"headcode" binding:"required"`               //请求头部结构
	Pathcode     string `json:"pathcode" binding:"required"`               //path结构
	Bodytype     string `json:"bodytype" binding:"required"`               //body的类型
	Bodyrecodes  string `json:"bodyrecodes" binding:"required"`            //body结构
	Returntype   string `json:"returntype" binding:"required"`             //返回类型
	Returncodes  string `json:"returncodes" binding:"required"`            //返回结构
	IsDeprecated int    `json:"isDeprecated" binding:"required,oneof=1 2"` //弃用状态
}

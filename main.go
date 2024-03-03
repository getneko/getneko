package main

import (
	"flag"
	_ "getneko/docs"
	"getneko/router"
	"strconv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           getneko api doc
// @version         1.0
// @host      localhost:8080

func main() {

	r := router.Router()
	//根据flag导入swagger依赖
	var enabledoc bool
	flag.BoolVar(&enabledoc, "doc", false, "enable swagger doc")
	//指定端口
	var flagports int
	flag.IntVar(&flagports, "port", 61223, "set server port")
	flag.Parse()
	runports := ":" + strconv.Itoa(flagports)
	if enabledoc {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	err := r.Run(runports)
	if err != nil {
		panic(err)
	}
}

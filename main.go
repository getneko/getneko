package main

import (
	_ "getneko/docs"
	"getneko/router"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           getneko api doc
// @version         1.0
// @host      localhost:8080

func main() {

	r := router.Router()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run()
	if err != nil {
		panic(err)
	}
}

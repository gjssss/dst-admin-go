package main

import (
	"dst-admin-go/bootstrap"
	"dst-admin-go/config/global"
	"dst-admin-go/router"
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	bootstrap.Init()
	app := router.NewRoute()
	app.Use(gzip.Gzip(gzip.BestCompression))
	app.Static("/static", "./web/static")

    // // 根路径 / 渲染 React 的 index.html 文件
    app.NoRoute(func(c *gin.Context) {
        c.File("./web/index.html")
    })
	err := app.Run(":" + global.Config.Port)
	if err != nil {
		fmt.Println("启动失败！！！", err)
	}

}

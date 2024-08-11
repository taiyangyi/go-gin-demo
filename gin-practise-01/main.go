package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// 创建一个默认的路由引擎
	r := gin.Default()
	// 配置路由
	// 当客户端以GET方法请求/ping路径时，会执行后面的匿名函数
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务,也可以指定端口号r.Run(":8000")
	r.Run()
}

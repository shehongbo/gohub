package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	//初始化 Gin 实例
	r := gin.Default()

	//注册一个路由
	r.GET("/", func(context *gin.Context) {
		// 以json 格式响应
		context.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})

	//处理404 请求
	r.NoRoute(func(context *gin.Context) {
		//获取表头信息的 Accept 信息
		acceptString := context.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			context.String(http.StatusNotFound, "页面返回404")
		} else {
			// 默认返回JSON
			context.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
	//运行服务
	r.Run(":8000")
}

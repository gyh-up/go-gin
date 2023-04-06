package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建服务
	ginService := gin.Default()
	
	ginService.GET("/hello", func(context *gin.Context) {
		
	})
}

package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建服务
	ginService := gin.Default()

	// session 使用
	store := cookie.NewStore([]byte("secret"))
	adminGroup := ginService.Group("/admin", sessions.Sessions("mysession", store))
	adminGroup.GET("getUserAuth", func(context *gin.Context) {
		session := sessions.Default(context)
		if session.Get("hello") == nil {
			session.Set("hello", "world")
			session.Save()
		}
		context.JSON(200, gin.H{
			"msg" : session.Get("hello"),
		})
	})




	// 服务端口
	ginService.Run(":8082")
}

func myHandler() (gin.HandlerFunc) {
	return func(context *gin.Context) {
		context.Set("session","123")
		session, _	 := context.Get("session")
		fmt.Println(session)
	}
}
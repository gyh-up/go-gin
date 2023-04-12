package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-gin/ginchat/docs"
	"go-gin/ginchat/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("index", service.GetIndex)
	r.GET("get_user_list", service.GetUserList)
	r.GET("create_user", service.CreateUser)
	r.GET("delete_user", service.DeleteUser)
	r.GET("update_user", service.UpdateUser)
	r.GET("SendMsg", service.SendMsg)

	return r
}

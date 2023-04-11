package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strconv"
)

func main() {
	// 创建服务
	ginService := gin.Default()
	sourceName := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", sourceName)
	if err != nil {
	}
	defer db.Close()

	ginService.GET("sql", func(context *gin.Context) {
		cname := context.Query("cname")
		cid, _ := strconv.Atoi(context.Query("cid"))
		course := Course {
			cid,
			cname,
		}
		db.Table("course").Create(course)
	})

	// 服务端口
	ginService.Run(":8081")
}

type Course struct {
	Cid int
	CNAMEAa string `gorm:"column:cname"`
}


package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func main() {
	// 创建服务
	ginService := gin.Default()
	surname1 := "root:root@tcp(127.0.0.1:3306)/test"
	DB := databaseConnect(surname1)
	ginService.GET("sql", func(context *gin.Context) {
		cname := context.Query("cname")
		cid, _ := strconv.Atoi(context.Query("cid"))
		course := Course {
			cid,
			cname,
		}
		insert(DB, course)
	})

	// 服务端口
	ginService.Run(":8080")
}
func databaseConnect (sourceName string) (DB *sql.DB) {
	DB, err := sql.Open("mysql", sourceName)
	if err != nil {
		fmt.Println(err)
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println(err)
	}
	return DB
}
type Course struct {
	Cid int
	Cname string
}
func insert(db *sql.DB, course Course) {
	db.Begin()
	sql, err := db.Prepare("insert into course (cname,cid) values (?,?)")
	if err != nil {
		fmt.Println(err)
	}
	res, _ := sql.Exec(course.Cname, course.Cid)
	fmt.Println(res)
}

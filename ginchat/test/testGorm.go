package main

import (
	"fmt"
	"go-gin/ginchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/ginchat?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		panic(any("failed to connect database"))
	}
	Message := &models.Message{}
	db.AutoMigrate(Message)
	db.AutoMigrate(&models.Contact{})
	db.AutoMigrate(&models.GroupBasic{})

	//db.Create(userBasic)
	fmt.Println(Message)
	//fmt.Println(db.First(userBasic, 1))
}

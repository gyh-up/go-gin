package models

import (
	"fmt"
	"go-gin/ginchat/utils"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	Id uint `gorm:"column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	Name string `gorm:"column:name" json:"name"`
	Password string `gorm:"column:password" json:"password"`
	Phone string `gorm:"column:phone" json:"phone" valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email string `gorm:"column:email" json:"email" valid:"email"`
	Salt string `gorm:"column:salt" json:"salt"`
	ClientIp string `gorm:"column:client_ip" json:"client_ip"`
	Identity string `gorm:"column:identity" json:"identity"`
	ClientPort string `gorm:"column:client_port" json:"client_port"`
	LoginTime time.Time `gorm:"column:login_time" json:"login_time"`
	HeartbeatTime time.Time `gorm:"column:heartbeat_time" json:"heartbeat_time"`
	LoginOutTime time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogout bool `gorm:"column:is_logout" json:"is_logout"`
	DeviceInfo string `gorm:"column:device_info" json:"device_info"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic  {
	data := make([]*UserBasic, 10)
	fmt.Println("------")
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
		fmt.Println("------")
	}
	return data
}

func CreateUser(user UserBasic) *gorm.DB  {
	return utils.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB  {
	return utils.DB.Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB  {
	return utils.DB.Where("id  = ? ", user.Id).Updates(UserBasic{Name: user.Name, Password: user.Password})
}

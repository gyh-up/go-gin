package models

import "time"

type GroupBasic struct {
	Id        uint      `gorm:"column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	Name string
	OwnerId uint //谁的关系
	Icon string
	Type int
	Desc string
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}


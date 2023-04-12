package models

import "time"

type Contact struct {
	Id        uint      `gorm:"column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	OwnerId uint //谁的关系
	TargetId uint
	Type int	// 对应类型
	Desc string
}

func (table *Contact) TableName() string {
	return "contact"
}


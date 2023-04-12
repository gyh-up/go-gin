package models

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"gopkg.in/fatih/set.v0"
	"time"
)

type Message struct {
	Id        uint      `gorm:"column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	FormId uint // 发送者
	TargetId uint //接收者
	Type string	// 消息类型 群聊 私聊 广播
	Media int	// 消息类型 文字 图片 音频
	Content string
	Pic string
	Url string
	Desc string
	Amount int // 其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}

type Node struct {
	Conn *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

// 映射关系
var clientMap  map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

func Chat(write http.ResponseWriter, requset *http.Request) {
	query := requset.URL.Query()
	userId := query.Get("userId")
	MsgType := query.Get("type")
	targetId := query.Get("targetId")
	context := query.Get("context")
	isvalida := true
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool{
			return isvalida
		},
	}).Upgrade(write, requset, nil)
	if err != nil{
		fmt.Println(err)
		return
	}
}

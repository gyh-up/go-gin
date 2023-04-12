package service

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go-gin/ginchat/models"
	"go-gin/ginchat/utils"
	"net/http"
	"strconv"
	"time"
)

// GetUserList
// @Tags 首页
// @Success 200 {string} json{"code","message"}
// @Router /get_user_list [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message" :data,
	})
}

func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	user.Email = c.Query("email")
	user.Phone = c.Query("phone")

	password := c.Query("password")
	repassword := c.Query("repassword")
	_ , err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(403, gin.H{
			"message" : "校验失败",
		})
	}

	if password != repassword {
		c.JSON(403, gin.H{
			"message" : "两次输入密码不一致",
		})
	}
	user.Password = password
	res := models.CreateUser(user)
	c.JSON(200, gin.H{
		"message" :res,
	})
}

func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.Id = uint(id)
	res := models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message" :res,
	})
}

func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.Id = uint(id)
	user.Email = c.Query("email")
	user.Phone = c.Query("phone")

	user.Password = c.Query("password")
	_ , err := govalidator.ValidateStruct(user)
	if err != nil {
		c.JSON(403, gin.H{
			"message" : err,
		})
	}
	res := models.UpdateUser(user)
	c.JSON(200, gin.H{
		"message" :res,
	})
}

// 防止跨域站点伪造请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context)  {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)

	 MsgHandler(ws, c)
}
func MsgHandler(ws *websocket.Conn, c *gin.Context)  {
	msg, err := utils.Subscribe(c, utils.PublishKey)
	if err != nil {
		fmt.Println(err)
	}
	tm := time.Now().Format("2006-01-02 15:04:05")
	m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
	err = ws.WriteMessage(1, []byte(m))
	if err != nil {
		fmt.Println(err)
	}
}
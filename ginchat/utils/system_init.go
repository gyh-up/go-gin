package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)
var (
	DB *gorm.DB
	Red *redis.Client
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("ginchat\\config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("config app", viper.Get("app"))
	//fmt.Println("config mysql", viper.Get("mysql"))
}

func InitMySQL() {
	// 自定义日志模板，打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel: logger.Info,
			Colorful: true,
		},
	)
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	//if err != nil {
	//	panic(any("failed to connect database"))
	//}
	//user := models.UserBasic{}
	//DB.Find(&user)
	//fmt.Println(DB)
}

func InitRedis() {
	Red  = redis.NewClient(&redis.Options{
		Addr: viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB: viper.GetInt("redis.DB"),
		PoolSize: viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConns"),
	})
	//pong, err := Red.Ping().Result()
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(pong)
	//}
}

const (
	PublishKey = "websocket"
)

// Publish 发布消息到REDIS
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	fmt.Println("发布消息到REDIS", msg)
	err = Red.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// Subscribe 订阅REDIS消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Red.Subscribe(ctx, channel)
	fmt.Println("Subscribe", ctx)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("订阅REDIS消息", msg)
	return msg.Payload, err
}
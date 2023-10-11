package model

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"im/config"
	"log"
	"strconv"
)

var (
	ImRedis *redis.Client
	ImDB    *gorm.DB
)

func InitRedis() {
	// 创建 Redis 客户端
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(config.AppConfig))
	if err != nil {
		log.Fatal(err)
		return
	}
	redishost := viper.Get("redis.host")
	redisPort := viper.Get("redis.port")
	addr := fmt.Sprintf("%s:%s", redishost.(string), strconv.Itoa(redisPort.(int)))

	fmt.Println(addr)

	ImRedis := redis.NewClient(&redis.Options{
		Addr:     addr, // Redis 服务器地址和端口
		Password: "",   // Redis 服务器密码
		DB:       0,    // Redis 数据库索引
	})

	ctx := context.Background()
	pong, err := ImRedis.Ping(ctx).Result()
	if err != nil {
		fmt.Println("连接 Redis 失败:", err)
		return
	}
	fmt.Println("连接 Redis 成功，收到 PONG 响应:", pong)
}

func InitDB() {
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(config.AppConfig))
	if err != nil {
		log.Fatal(err)
		return
	}
	mysqlConf := viper.Get("mysql.dsn")
	fmt.Println(mysqlConf)
	ImDB, err = gorm.Open(mysql.Open(mysqlConf.(string)))

	if err != nil {
		fmt.Println(err)
	}
	AutoMigrateTable()
	fmt.Println("连接mysql成功！")
}

func InitKafka() {
	// 设置 Kafka 连接配置
	config := sarama.NewConfig()
	config.ClientID = "your-client-id"

	// 连接到 Kafka
	brokers := []string{"localhost:9092"}
	client, err := sarama.NewClient(brokers, config)
	if err != nil {
		fmt.Println("Failed to connect to Kafka:", err)
		return
	}
	client.Close()
}

func GetMysqlDB() *gorm.DB {
	return ImDB
}

func GetRedis() *redis.Client {
	return ImRedis
}

func AutoMigrateTable() {
	_ = ImDB.AutoMigrate(&User{})
}

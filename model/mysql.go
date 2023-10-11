package model

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ImDB *gorm.DB

func SB( ) {}




func InitDB() {
	dsn := viper.GetString("mysql.dsn")
	ImDB, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return
}

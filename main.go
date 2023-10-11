package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"im/config"
	"im/model"
)

func main() {
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(config.AppConfig))
	mysql_conf := viper.Get("mysql")

	model.ImDB.Row()
	fmt.Println(mysql_conf)
}

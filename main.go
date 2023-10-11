package main

import (
	"fmt"
	"im/model"
)

func main() {

	model.InitDB()

	db := model.GetMysqlDB()

	// 获取当前数据库的名称
	dbName := db.Migrator().CurrentDatabase()
	fmt.Println("当前数据库名称:", dbName)

	model.InitRedis()
	user := &model.User{
		UserName: "傻逼",
		PassWord: "asadaf",
	}

	db.Create(&user)
	//tables, err := db.DBMetas()

}

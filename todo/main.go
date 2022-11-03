package main

import (
	"todo/dao"
	"todo/models"
	"todo/router"
)

func main() {
	// 创建数据库
	// SQL：create databases bubble
	//链接MySQL数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	dao.DB.AutoMigrate(&models.Todo{})

	// 启动服务
	r := router.Router()

	err = r.Run()
	if err != nil {
		panic(err)
	}
}

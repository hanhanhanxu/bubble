package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
)

func main() {
	//1、连接数据库
	err := dao.InitDb()
	if err != nil {
		panic(err)
	}
	//现在不需要手动关闭了，都是使用的连接池

	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	//2、开启gin
	r := routers.SetupRouter()
	r.Run()
}

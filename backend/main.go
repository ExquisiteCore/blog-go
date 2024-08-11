package main

import (
	"backend/config"
	"backend/model"
	"backend/routes"
)

func init() {

}

func main() {
	config.Init()      //配置文件初始化
	model.InitDb()     //数据库初始化
	routes.InitRuter() //路由初始化
}

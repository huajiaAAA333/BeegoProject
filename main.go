package main

import (
	_ "BeeGoDemo/routers"
	"astaxie/beego"
	"fmt"
)

func main() {
	//定义config变量,接受并赋值为全局变量
	config := beego.AppConfig
	//获取配置变量
	appName := config.String("appname")
	fmt.Println("项目应用名称:",appName)
	port,err := config.Int("httpport")
	if err !=nil {
		//配置信息解析错误
		panic("项目配置信息错误,")

	}
	fmt.Println("应用监听端口:",port)

	driver := config.String("db_driver")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_passward")
	dbIP := config.String("db_IP")
	db




}


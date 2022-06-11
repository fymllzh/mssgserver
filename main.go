package main

import (
	"fmt"
	"mssgserver/config"
)

func main()  {
	//测试调用
	config.Test()
	//获取配置文件的值
	fmt.Println(config.File.MustValue("login_server","host","128.0.0.2"))
}

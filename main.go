package main

import (
	"fmt"
	"mssgserver/server"
)
func main()  {
	////测试调用长链接
	//config.Test()
	////获取配置文件的值
	//fmt.Println(config.File.MustValue("login_server","host","128.0.0.2"))

	//获取配置信息
	//str :=utils.Viper.GetString("db.Host")
	//fmt.Println(str)


	var u server.User

	//插入数据
	//err := u.InsertRowDemo()
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//单个查询
	//err := u.SelectRowDemo(2)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Printf("%+v",u)

	//查询列表
	res := u.SelectListDemo()
	fmt.Println(len(res))



}

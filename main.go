package main

import (
	"fmt"
	"mssgserver/utils"
)

type User struct {
	ID   int64
	Name string
}
func main()  {
	////测试调用长链接
	//config.Test()
	////获取配置文件的值
	//fmt.Println(config.File.MustValue("login_server","host","128.0.0.2"))

	//获取配置信息
	//str :=utils.Viper.GetString("db.Host")
	//fmt.Println(str)


	//var u server.User

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
	//res := u.SelectListDemo()
	//fmt.Println(len(res))

	//更新
	//num :=u.UpdateRowDemo(0,101)
	//fmt.Println(num)

	//删除
	//b :=u.DeleteRowDemo(2)
	//fmt.Println(b)

	//测试redis 需要结构体存取
	//u := User{
	//	ID : 9,
	//	Name: "nihao",
	//
	//}
	//utils.SetJson("haha",u,3600)
	var aa User
	if utils.GetJson("haha",&aa) {
		fmt.Printf("%+v",aa)
	} else {
		fmt.Println(888)
	}


	//redis结构体存取
	//utils.StructAdd()
	//utils.StructValues()


}

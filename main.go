package main

type User struct {
	ID   int64
	Name string
}
func main()  {
	////测试调用长链接
	//config.Test()
	////获取配置文件的值
	//fmt.Println(config.File.MustValue("login_server","host","128.0.0.2"))
///////////////////////读取配置信息//////////////////////
	//获取配置信息
	//str :=utils.Viper.GetString("db.Host")
	//fmt.Println(str)

	///////////////////////CURD//////////////////////
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

	///////////////////////////redis////////////////
	//测试redis 需要结构体存取
	//u := User{
	//	ID : 9,
	//	Name: "nihao",
	//
	//}
	//utils.SetJson("haha",u,3600)
	//var aa User
	//if utils.GetJson("haha",&aa) {
	//	fmt.Printf("%+v",aa)
	//} else {
	//	fmt.Println(888)
	//}

	//redis 读取字符串
	//conn :=utils.Rdb.Get()
	//defer conn.Close()
	//res,_ :=conn.Do("GET","ceshikey") //可以换成别的命令
	//fmt.Println(string(res.([]byte)))


	//redis结构体存取
	//utils.StructAdd()
	//utils.StructValues()

////////////////////gin单一设置web/////////////////////
//	r := gin.Default()
//	r.GET("/hello", helloHandler)
//	if err := r.Run("127.0.0.1:8005"); err != nil {
//		fmt.Println("startup service failed, err:%v\n", err)
//	}
//}
//
//func helloHandler(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{
//		"message": "Hello q1mi!",
//	})
	////////////////gin封装route//////////////
	//r := gin.Default()
	//router.InitRouter(r)
	//if err := r.Run(":8005"); err != nil {
	//	fmt.Println("启动错误")
	//}

	////////测试时间包//////////
	//str := utils.GetDateTime()
	//fmt.Println(str)
	//unix := utils.GetUnix()
	//fmt.Println(unix)
	//date := utils.Unix2DateTime(unix)
	//fmt.Println(date)
	//u := utils.Date2Unix(str)
	//fmt.Println(u)



}

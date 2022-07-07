package main

import (
	"flag"
	"fmt"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"mssgserver/api"
	"mssgserver/router"
	"mssgserver/utils"
	"html/template"
)

type User struct {
	ID   int64
	Name string
}
func main()  {
	//异常捕获只能捕获单前文件的异常
	defer func() {
		if err := recover(); err != nil {
			utils.Logger.Info("捕获异常错误")
			if str, ok := err.(string); ok {
				utils.Logger.Info("异常是 " + str)
			}
		}
	}()

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
	r := gin.Default()
	//测试flag包--初始化日志目录
	var logPath string
	flag.StringVar(&logPath,"z","logs","这是一个日志目录")
	flag.Parse()
	utils.InitLog(logPath)
	//初始化数据库
	utils.InitDb()

	//初始化路由
	router.InitRouter(r)

	//引入static静态资源
	r.Static("/statics","web/static")

	//可以使用中间件
	//loginTpl := gintemplate.NewMiddleware(gintemplate.TemplateConfig{
	//	Root:         "web/templates/login",
	//	Extension:    ".html",
	//	Master:       "",
	//	Partials:     []string{},
	//	DisableCache: true,
	//})
	//login := r.Group("/admin",loginTpl)

	//不使用中间件
	r.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "web/templates/login",
		Extension:    ".html",
		Master:       "",
		DisableCache: true,
		Funcs:        template.FuncMap{},
	})

	//登录页面管理
	login := r.Group("/admin")
	{
		login.GET("/login", api.Login)
		login.POST("/logins", api.LoginAuth)
	}

	if err := r.Run(":8005"); err != nil {
		fmt.Println("启动错误")
	}

	////////测试时间包//////////
	//str := utils.GetDateTime()
	//fmt.Println(str)
	//unix := utils.GetUnix()
	//fmt.Println(unix)
	//date := utils.Unix2DateTime(unix)
	//fmt.Println(date)
	//u := utils.Date2Unix(str)
	//fmt.Println(u)

	//测试jwt
	//token,err := utils.GenToken("li","zhi",6000)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//	fmt.Println(token)
	//
	//var aa *utils.Claims
	//	aa,err = utils.ValidToken(token,"zhi")
	//	if err != nil {
	//		fmt.Println(err)
	//	}

		//fmt.Println(aa.Data)


}

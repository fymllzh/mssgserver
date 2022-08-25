package main

import (
	"fmt"
	"mssgserver/config"
	"mssgserver/net"
)

func main()  {
	host := config.File.MustValue("login_server","host")
	port := config.File.MustValue("login_server","port")

	s := net.NewServer(host + ":" + port)

	s.Start()
	fmt.Println("登陆成功了啊")
}
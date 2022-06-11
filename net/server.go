package net

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

type server struct {
	addr string
	route *route
}

func NewServer(addr string) *server {
	return &server{
		addr : addr,
	}
}

func (s *server) Start() {
	http.HandleFunc("/",s.wsHandler)
	err := http.ListenAndServe(s.addr,nil)
	if err != nil {
		panic(err)
	}
}

// http升级websocket协议的配置
var wsUpgrader = websocket.Upgrader{
	// 允许所有CORS跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *server) wsHandler(w http.ResponseWriter, r *http.Request)  {
	//思考websocket
	wsConn,err := wsUpgrader.Upgrade(w,r,nil)
	if err != nil {
		log.Fatalln("ws 链接失败",err)
	}
	log.Println("ws 链接成功")

	//响应客户端
	err = wsConn.WriteMessage(websocket.BinaryMessage,[]byte("hello"))
	fmt.Println(err)


	//读取客户端的信息
	for {
		_,str,err := wsConn.ReadMessage()
		if err != nil {
			log.Fatalln("ws 读取",err)
		}
		fmt.Println(string(str))
	}


}



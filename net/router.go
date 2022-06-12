package net

type HandlerFunc func(name string, req *WsMsgReq, rsp *WsMsgRsp)

type group struct {
	prefix string
	handlerMap map[string]HandlerFunc
}

type Router struct {
	group []*group
}
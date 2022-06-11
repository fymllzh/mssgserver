package net

type HandlerFunc func()

type group struct {
	prefix string
	handlerMap map[string]HandlerFunc
}

type route struct {
	group []*group
}
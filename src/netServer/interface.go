package netServer

import "sync"

// 实体的定义

type NetServer struct {
	port      uint32
	netRouter NetRouter
}

type NetParam struct {
	Values map[string]string
}

type NetRouter struct {
	routerList map[string]func(*NetParam, *NetParam) string
	routerLock sync.Mutex
}

type RouterHandler struct {

}

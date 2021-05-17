package session

import "fmt"

// 中间件让用户选择使用哪个版本
var (
	sessionMgr SessionMgr
)

func Init(provider, addr string, options ...string) {
	switch provider {
	case "memory":
		sessionMgr = NewMemorySessionMgr()
	case "redis":
		sessionMgr = NewRedisSessionMgr()
	default:
		fmt.Errorf("不支持")
		return
	}
	sessionMgr.Init(addr, options...)
	return

}

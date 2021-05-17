package session

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"sync"
)

// MemorySessionMgr设计
// 定义MemorySessionMgr对象(字段:存放所有session的map,读写锁)
// 构造函数
// Init()
// CreateSession()
// GetSession()

// MemorySessionMgr 定义对象
type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwLock     sync.RWMutex
}

func NewMemorySessionMgr() SessionMgr {
	sr := &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
	return sr
}

func (s *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}

func (s *MemorySessionMgr) CreateSession() (session Session, err error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	// go get github.com/satori/go.uuid
	var id uuid.UUID
	id = uuid.NewV4()
	// 转string
	sessionId := id.String()
	session = NewMemorySession(sessionId)
	// 加入到大map
	s.sessionMap[sessionId] = session
	return
}

func (s *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	session, ok := s.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return
}

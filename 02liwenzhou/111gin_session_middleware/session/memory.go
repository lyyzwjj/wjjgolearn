package session

import (
	"errors"
	"sync"
)

// MemorySession设计:
// 定义MemorySession对象(字段:sessionId\存kv的map,读写锁)
// 构造函数, 为了获取对象
// Set()
// Get()
// Del()
// Save()

// MemorySession 具体实现 具体对象
type MemorySession struct {
	sessionId string
	// 存kv
	data   map[string]interface{}
	rwLock sync.RWMutex
}

// NewMemorySession 构造函数
func NewMemorySession(id string) *MemorySession {
	s := &MemorySession{
		sessionId: id,
		data:      make(map[string]interface{}, 16),
	}
	return s
}

func (m *MemorySession) Set(key string, value interface{}) (err error) {
	// 加锁
	m.rwLock.Lock()
	// 解锁
	defer m.rwLock.Unlock()
	// 设置值
	m.data[key] = value
	return
}

func (m *MemorySession) Get(key string) (value interface{}, err error) {
	// 加锁
	m.rwLock.Lock()
	// 解锁
	defer m.rwLock.Unlock()
	// 获取值
	value, ok := m.data[key]
	if !ok {
		err = errors.New("key not exists in session")
		return
	}
	return
}

func (m *MemorySession) Del(key string) (err error) {
	// 加锁
	m.rwLock.Lock()
	// 解锁
	defer m.rwLock.Unlock()
	delete(m.data, key)
	return
}

func (m *MemorySession) Save() (err error) {
	return
}

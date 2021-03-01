package maps

import "sync"

/**
 * @author  wjj
 * @date  2020/9/13 9:35 下午
 * @description
 */
func CreateSyncMapBenchmarkAdapter() *SyncMapBenchmarkAdapter {
	return &SyncMapBenchmarkAdapter{}
}

type SyncMapBenchmarkAdapter struct {
	m sync.Map
}

func (m *SyncMapBenchmarkAdapter) Set(key interface{}, val interface{}) {
	m.m.Store(key, val)
}

func (m *SyncMapBenchmarkAdapter) Get(key interface{}) (interface{}, bool) {
	return m.m.Load(key)
}

func (m *SyncMapBenchmarkAdapter) Del(key interface{}) {
	m.m.Delete(key)
}
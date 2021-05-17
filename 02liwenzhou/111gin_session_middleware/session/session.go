package session

type Session interface {
	// Set 将session存到内存中的map
	Set(key string, value interface{}) error
	// Get 取数据,实现延迟加载
	Get(key string) (interface{}, error)
	Del(key string) error
	Save() error
}

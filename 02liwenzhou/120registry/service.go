package registry

// Service 抽象服务
type Service struct {
	// 服务名
	Name string `json:"name"`
	// 节点列表
	Nodes []*Node `json:"nodes"`
}

// Node 单个服务节点的抽象
type Node struct {
	Id     string `json:"id"`
	Ip     string `json:"ip"`
	Port   int    `json:"port"`
	Weight int    `json:"weight"`
}

// AllServiceInfo 存放所有服务信息 存入atomic.Value 为了防止并发
type AllServiceInfo struct {
	ServiceMap map[string]*Service //节点信息 servicename node
}

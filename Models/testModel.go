package Models

// 每个监控项抽象类
type Item interface {
	// 每个监控项的实例id
	GetID() string
	// 每个监控项的名称
	GetName() string
	// 每个监控项的最新数值
	GetValue() interface{}
	// 数值的支持类型 例如 int string float bool 等类型
	GetType() string
	// 每个监控项的models属性,用于映射cmdb使用
	GetModels() interface{}
	// 每个监控项目采集频率 单位秒
	GetRate() int
	// 属性每个监控项的属性
	GetTAG() string
}

/*
Host 节点
*/
type Host []Item

// 一个监控主机有多少监控项
func (nodes Host) Len() int {
	return len(nodes)
}

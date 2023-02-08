package observer

// 主体
type Subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}

package Singleton

type eagerSingleton struct {
	// 将私有的 instance 变量放在结构体中
}

var eagerInstance *eagerSingleton = &eagerSingleton{}

func GetEagerInstance() *eagerSingleton {
	return eagerInstance
}

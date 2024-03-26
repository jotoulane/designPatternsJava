package Singleton

import "sync"

type lazySingleton struct {
	// 将私有的 instance 变量放在结构体中
}

var lazyInstance *lazySingleton
var once sync.Once

func GetLazyInstance() *lazySingleton {
	once.Do(func() {
		lazyInstance = &lazySingleton{}
	})
	return lazyInstance
}

package singleton

import "sync"

//1. 传统思想方式实现单例模式

var (
	ins  *instance
	lock sync.Mutex
)

// Instance 单例私有化
type instance struct {
}

// GetInstance 双重检查 这种方式其实很java的一样
func GetInstance() *instance {
	if ins == nil {
		lock.Lock()
		defer lock.Unlock()
		if ins == nil {
			ins = &instance{}
		}
	}
	return ins
}

//2. go的特有方式

type singleton struct{}

var (
	sn   *singleton
	once sync.Once
)

// GetInstance_ 使用go特有的方式实现单例模式
func GetInstance_() *singleton {
	once.Do(func() {
		sn = &singleton{}
	})
	return sn
}

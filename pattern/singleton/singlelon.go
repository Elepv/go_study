package singleton

import "sync"

//Singleton 是单例模式类
type Singleton struct {
	//实现业务逻辑

}

//定义singleton包中的私有变量
var singleton *Singleton
var once sync.Once

//GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}

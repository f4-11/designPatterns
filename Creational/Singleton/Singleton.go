package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type singleton struct{}

var initialized uint32

var instance *singleton // = new(singleton)饿汉式，编译就有

var lock sync.Mutex

func GetInstance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = new(singleton) //懒汉式，可能导致并发时创建多个对象,所以要加锁.单例模式创建全局唯一的变量
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

/*
也可以使用下面的代码替代：
*
var Once sync.Once

	func Getinstance() *instance {
		once.do(func(){
		instance = new(singleton)
		})
		return instance
	}
*/
func (s *singleton) fucku() {
	fmt.Println("fucku")
}

func main() {
	GetInstance().fucku()
}

package singleton

import (
	"fmt"
	"sync"
)

type SingleLock struct {
}

var lock sync.Mutex
var lockSingle *SingleLock

func getInstanceLock() *SingleLock {
	if lockSingle == nil {
		lock.Lock()
		defer lock.Unlock()
		//在获取到锁后还会有另一个 nil检查。 这是为了确保即便是有多个协程绕过了第一次检查，
		//也只能有一个可以创建单例实例。 否则， 所有协程都会创建自己的单例结构体实例
		if lockSingle == nil {
			lockSingle = &SingleLock{}
			fmt.Println("lockSingle has been created")
		} else {
			fmt.Println("lockSingle is already been created1")
		}
	} else {
		fmt.Println("lockSingle is already been created2")
	}

	return lockSingle
}

//单例的各种实现
package singleton


import (
	"sync"
	"sync/atomic"
)

type singleton struct {

}

var sg *singleton

//simple
func GetSingleton() *singleton  {
	if sg == nil{
		sg = &singleton{}
	}
	return sg
}

//with lock
var mu sync.Mutex
func GetSingletonWithLock() *singleton {
	mu.Lock()
	defer mu.Unlock()
	if sg == nil{
		sg = &singleton{}
	}
	return sg
}

//with check
func GetSingletonWithCheck() *singleton {
	if sg == nil {
		mu.Lock()
		defer mu.Unlock()
		if sg == nil{
			sg = &singleton{}
		}
	}
	return sg
}

//with Atomic
var initData uint32
func GetSingletonWithAtomic() *singleton {
	if atomic.LoadUint32(&initData) == 1 {
		return sg
	}
	mu.Lock()
	defer mu.Unlock()
	if initData == 0 {
		sg = &singleton{}
		atomic.StoreUint32(&initData, 1)
	}
	return sg
}

//with Once
var once sync.Once
func GetSingletonWithOnce() *singleton {
	once.Do(func() {
		sg = &singleton{}
	})
	return sg
}
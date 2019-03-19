package singleton

import (
	"sync"
	"sync/atomic"
)

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

var mux sync.Mutex

/*
This is a better approach, but still is not perfect.
Since due to compiler optimizations there is not an atomic check on the instance store state.
http://marcio.io/2015/07/singleton-pattern-in-go/
 */
func GetInstanceCheckLockCheck() *singleton {
	if instance == nil {
		mux.Lock()
		defer mux.Unlock()
		if instance == nil {
			instance = new(singleton)
		}
	}
	return instance
}

/*
http://marcio.io/2015/07/singleton-pattern-in-go/
 */
var initialized uint32
func GetInstanceAtomicCheckLockCheck() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mux.Lock()
	defer mux.Unlock()
	if initialized == 0 {
		instance = &singleton{}
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}

var once sync.Once
func GetInstanceOnce() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
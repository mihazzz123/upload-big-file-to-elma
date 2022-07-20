package syncutils

import (
	"sync"
	"sync/atomic"
)

// OnceWithReset реализация sync.Once с возможностью возобновления действия
type OnceWithReset struct {
	done uint32
	m    sync.Mutex
}

// Do выполнить переданную функцию
func (o *OnceWithReset) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	// Slow-path.
	o.m.Lock()
	defer o.m.Unlock()
	if atomic.LoadUint32(&o.done) == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

// Reset установить OnceWithReset в невыполненное состояние
func (o *OnceWithReset) Reset() {
	o.m.Lock()
	defer o.m.Unlock()
	o.reset()
}

func (o *OnceWithReset) reset() {
	atomic.StoreUint32(&o.done, 0)
}

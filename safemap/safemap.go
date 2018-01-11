package safemap

import "sync"

type safeMap struct {
	sync.RWMutex
	items map[string]interface{}
}

func New() *safeMap {
	return &safeMap{items: make(map[string]interface{})}
}

func (sm *safeMap) Set(k string, v interface{}) {
	sm.Lock()
	defer sm.Unlock()

	sm.items[k] = v
}

func (sm *safeMap) Get(k string) (interface{}, bool) {
	sm.RLock()
	defer sm.RUnlock()

	v, exist := sm.items[k]
	return v, exist
}

func (sm *safeMap) Size() int {
	sm.RLock()
	defer sm.RUnlock()

	return len(sm.items)
}

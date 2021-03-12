package waitingmap

import "sync"

type OnceMap struct {
	store sync.Map
}

func (m *OnceMap) Do(key string, f func()) bool {
	return m.DoInterface(key, f)
}
func (m *OnceMap) DoInterface(k interface{}, f func()) bool {
	_, loaded := m.store.LoadOrStore(k, true)
	if loaded {
		return false
	}
	defer m.store.Delete(k)
	f()
	return true
}

func NewOnceMap() *OnceMap {
	return &OnceMap{}
}

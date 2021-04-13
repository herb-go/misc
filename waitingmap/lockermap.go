package waitingmap

import "sync"

type lockerentity struct {
	count   int64
	rwcount int64
	sync.RWMutex
}

type locker struct {
	m   *LockerMap
	key interface{}
}

func (l *locker) Lock() {
	l.m.LockInterface(l.key)
}

func (l *locker) Unlock() {
	l.m.UnlockInterface(l.key)
}

type rlocker struct {
	m   *LockerMap
	key interface{}
}

func (l *rlocker) Lock() {
	l.m.RLockInterface(l.key)
}
func (l *rlocker) Unlock() {
	l.m.RUnlockInterface(l.key)
}

type LockerMap struct {
	locker sync.Mutex
	store  map[interface{}]*lockerentity
}

func (m *LockerMap) Lock(key string) bool {
	return m.LockInterface(key)
}
func (m *LockerMap) LockInterface(k interface{}) bool {
	var newlock bool
	m.locker.Lock()
	l := m.store[k]
	if l == nil {
		l = &lockerentity{}
		m.store[k] = l
		newlock = true
	}
	l.count++
	m.locker.Unlock()
	l.Lock()
	return newlock
}
func (m *LockerMap) Unlock(key string) bool {
	return m.UnlockInterface(key)
}
func (m *LockerMap) UnlockInterface(k interface{}) bool {
	var free bool
	m.locker.Lock()
	l := m.store[k]
	l.count--
	if l.count == 0 && l.rwcount == 0 {
		delete(m.store, k)
		free = true
	}
	m.locker.Unlock()
	l.Unlock()
	return free
}

func (m *LockerMap) RLock(key string) bool {
	return m.RLockInterface(key)
}
func (m *LockerMap) RLockInterface(k interface{}) bool {
	var newlock bool
	m.locker.Lock()
	l := m.store[k]
	if l == nil {
		l = &lockerentity{}
		m.store[k] = l
		newlock = true
	}
	l.rwcount++
	m.locker.Unlock()
	l.RLock()
	return newlock
}
func (m *LockerMap) RUnlock(key string) bool {
	return m.RUnlockInterface(key)
}
func (m *LockerMap) RUnlockInterface(k interface{}) bool {
	var free bool
	m.locker.Lock()
	l := m.store[k]
	l.rwcount--
	if l.count == 0 && l.rwcount == 0 {
		delete(m.store, k)
		free = true
	}
	m.locker.Unlock()
	l.RUnlock()
	return free
}

func (m *LockerMap) Locker(key string) sync.Locker {
	return m.LockerInterface(key)
}
func (m *LockerMap) LockerInterface(k interface{}) sync.Locker {
	return &locker{
		m:   m,
		key: k,
	}
}
func (m *LockerMap) RLocker(key string) sync.Locker {
	return m.RLockerInterface(key)
}
func (m *LockerMap) RLockerInterface(k interface{}) sync.Locker {
	return &rlocker{
		m:   m,
		key: k,
	}
}
func NewLockerMap() *LockerMap {
	return &LockerMap{
		store: map[interface{}]*lockerentity{},
	}
}

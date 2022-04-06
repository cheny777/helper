package util

import "sync"

func NewSafeMap() *SafeMap {
	safeMap := &SafeMap{
		innerMap: make(map[interface{}]interface{}),
	}
	return safeMap
}

type SafeMap struct {
	innerMap map[interface{}]interface{}
	lock     sync.Mutex
}

//如果key存在则不操作
func (m *SafeMap) Insert(key interface{}, val interface{}) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, ok := m.innerMap[key]; !ok {
		m.innerMap[key] = val
		return true
	}
	return false
}

func (m *SafeMap) Set(key interface{}, val interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.innerMap[key] = val
	return
}

func (m *SafeMap) Exist(key interface{}) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	_, ok := m.innerMap[key]
	return ok
}

func (m *SafeMap) Get(key interface{}) (interface{}, bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	val, ok := m.innerMap[key]
	return val, ok
}

func (m *SafeMap) GetAll() []interface{} {
	m.lock.Lock()
	defer m.lock.Unlock()
	var list []interface{}
	for _, val := range m.innerMap {
		list = append(list, val)
	}
	return list
}

func (m *SafeMap) Delete(key interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.innerMap, key)
	return
}

func (m *SafeMap) Length() int {
	m.lock.Lock()
	defer m.lock.Unlock()
	return len(m.innerMap)
}

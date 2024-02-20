package basic

import "sync"

type SyncMap[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

func (m *SyncMap[K, V]) Get(key K) V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.m[key]
}

func (m *SyncMap[K, V]) Values() []V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := []V{}
	for _, v := range m.m {
		result = append(result, v)
	}
	return result
}

func (m *SyncMap[K, V]) Set(key K, value V) {
	m.mu.Lock()
	m.m[key] = value
	m.mu.Unlock()
}

func (m *SyncMap[K, V]) Delete(key K) {
	m.mu.Lock()
	delete(m.m, key)
	m.mu.Unlock()
}

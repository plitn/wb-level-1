package task7

import "sync"

/*
по сути обычная мапа с "привязанным" мьютексом
просто когда читаем и записываем в мапу мы локаем мьютекс
*/
type safeMap struct {
	mu sync.Mutex
	m  map[int]int
}

func (m *safeMap) read(k int) (value int, ok bool) {
	m.mu.Lock()
	value, ok = m.m[k]
	m.mu.Unlock()
	return value, ok
}
func (m *safeMap) write(k int, v int) {
	m.mu.Lock()
	m.m[k] = v
	m.mu.Unlock()
}

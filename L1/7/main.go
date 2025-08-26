package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mutex sync.RWMutex
	data  map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{data: make(map[string]int)}
}

func (s *SafeMap) Put(key string, value int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data[key] = value
}

func (s *SafeMap) Get(key string) int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.data[key]
}

func main() {
	safeMap := NewSafeMap()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				safeMap.Put(fmt.Sprintf("key_%d", i), i)
			}
		}(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				safeMap.Get(fmt.Sprintf("key_%d", i))
			}
		}()
	}
	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]int),
	}
}

// увеличивает значение по ключу на 1
func (s *SafeMap) IncKey(key string) {
	s.mu.Lock()
	s.m[key]++
	s.mu.Unlock()
}

// записывает значение по ключу
func (s *SafeMap) Set(key string, v int) {
	s.mu.Lock()
	s.m[key] = v
	s.mu.Unlock()
}

// возвращает значение и флаг наличия ключа
func (s *SafeMap) Get(key string) (int, bool) {
	s.mu.Lock()
	v, ok := s.m[key]
	s.mu.Unlock()
	return v, ok
}

// возвращает копию map
func (s *SafeMap) Snapshot() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make(map[string]int, len(s.m))
	for k, v := range s.m {
		out[k] = v
	}
	return out
}

func main() {
	const goroutines = 100
	const perGor = 1000

	sm := NewSafeMap()
	sm.Set("counter", 10)
	wg := &sync.WaitGroup{}
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < perGor; j++ {
				sm.IncKey("counter")
			}
		}()
	}
	wg.Wait()
	snap := sm.Snapshot()
	if v, ok := sm.Get("counter"); ok {
		fmt.Printf("Final value (Get): %d\n", v)
	}
	fmt.Println("Final counts:", snap)
}

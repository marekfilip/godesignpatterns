// Package singleton is a example implementation of singleton design pattern in Go
package singleton

// Singleton interface exposes methods but keeps
// instance object hidden and its properties
type Singleton interface {
	AddOne() int
}

type singleton struct {
	counter int
}

var instance *singleton

// GetInstance initialize Singleton (if not initialized)
// and returns its instance
func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}

	return instance
}

// AddOne increments Singleton counter
func (s *singleton) AddOne() int {
	s.counter++
	return s.counter
}

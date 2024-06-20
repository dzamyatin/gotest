package singleton

import (
	"reflect"
	"sync"
)

var (
	s        sync.RWMutex
	services = make(map[string]interface{})
)

func GlobalGetOrCreateTyped[T interface{}](create func() T) T {
	var t T
	return GlobalGetOrCreate(reflect.TypeOf(t).String(), create)
}

func GlobalGetOrCreate[T interface{}](name string, create func() T) T {
	s.RLock()
	if res, ok := services[name]; ok {
		s.RUnlock()
		return res.(T)
	}
	s.RUnlock()

	res := create()

	s.Lock()
	services[name] = res
	s.Unlock()

	return res
}

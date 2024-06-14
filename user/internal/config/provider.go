package config

import (
	"reflect"
	"sync"
)

var (
	s        sync.RWMutex
	services = make(map[string]interface{})
)

func syncGetOrCreateByType[T interface{}](create func() T) T {
	var t T
	return syncGetOrCreate(reflect.TypeOf(t).String(), create)
}

func syncGetOrCreate[T interface{}](name string, create func() T) T {
	s.RLock()
	if res, ok := services[name]; ok {
		return res.(T)
	}
	s.RUnlock()

	res := create()
	s.Lock()
	services[name] = res
	s.Unlock()

	return services[name].(T)
}

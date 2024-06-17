package config

import (
	"context"
	"reflect"
	"sync"
)

type Session struct {
	mx       sync.RWMutex
	services map[string]interface{}
	ctx      context.Context
}

func NewSession(ctx context.Context) *Session {
	return &Session{
		ctx: ctx,
	}
}

func getOrCreateTyped[T interface{}](s *Session, create func() T) T {
	var t T
	return getOrCreate(s, reflect.TypeOf(t).String(), create)
}

func getOrCreate[T interface{}](s *Session, name string, create func() T) T {
	s.mx.RLock()
	if res, ok := services[name]; ok {
		s.mx.RUnlock()
		return res.(T)
	}
	s.mx.RUnlock()

	res := create()

	s.mx.Lock()
	services[name] = res
	s.mx.Unlock()

	return res
}

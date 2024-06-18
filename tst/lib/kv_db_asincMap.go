package lib

import (
	"sync"
	"time"
)

type KVDBAsyncMap struct {
	mu sync.RWMutex
	m  sync.Map
	db KVDatabase
}

func NewKVDBAsyncMap(db KVDatabase) KVDBAsyncMap {
	return KVDBAsyncMap{db: db}
}

func (k *KVDBAsyncMap) Get(key string) (string, error) {
	k.mu.RLock()
	if v, ok := k.m.Load(key); ok {
		res, _ := v.(cacheUnit)
		k.mu.RUnlock()
		return res.v, res.err
	}
	k.mu.RUnlock()
	k.mu.Lock()
	defer k.mu.Unlock()

	v, err := k.db.Get(key)
	k.m.Store(key, cacheUnit{
		v:   v,
		err: err,
		tta: time.Now().Unix(),
	})

	return v, err
}

func (k *KVDBAsyncMap) Set(key, value string) error {
	return k.db.Set(key, value)
}

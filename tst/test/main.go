package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	db := NewKVDB()
	cache := NewKVCache(db)

	for i := 0; i < 100; i++ {
		go func(i int) {
			cache.Set("Hella", strconv.Itoa(i))
		}(i)
	}

	strt := time.Now()
	for i := 0; i < 1000000; i++ {
		//for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			cache.Get("Hella")
			//fmt.Println(
			//	cache.Get("Hella"),
			//)
			wg.Done()
		}()
	}

	wg.Wait()

	log.Println("Met:")
	log.Println(metric.Load())
	end := time.Now().Sub(strt)
	fmt.Printf("Time: %s", end)
}

var metric = atomic.Int32{}

var (
	KeyNotFound = errors.New("Key not found")
)

type KVDatabase interface {
	Get(key string) (string, error)
	Set(key, value string) error
}

const TTL_CACHE_SEC = 100

type cacheUnit struct {
	v   string
	err error
	tta int64
}

type KVCache struct {
	mu  sync.RWMutex
	m   map[string]cacheUnit
	db  KVDatabase
	mxx sync.RWMutex
	mx  map[string]*sync.RWMutex
}

func NewKVCache(db KVDatabase) *KVCache {
	return &KVCache{
		m:  make(map[string]cacheUnit),
		db: db,
		mx: make(map[string]*sync.RWMutex),
	}
}

func (k *KVCache) Get(key string) (string, error) {
	k.mu.RLock()

	if v, ok := k.m[key]; ok {
		k.mu.RUnlock()
		k.invalidate(key, v)

		return v.v, v.err
	}

	k.mu.RUnlock()

	return k.update(key)
}

func (k *KVCache) getKeyMx(key string) *sync.RWMutex {
	k.mxx.RLock()
	if mx, ok := k.mx[key]; ok {
		k.mxx.RUnlock()
		return mx
	}
	k.mxx.RUnlock()

	k.mxx.Lock()
	defer k.mxx.Unlock()

	mx := &sync.RWMutex{}
	k.mx[key] = mx

	return mx
}

func (k *KVCache) invalidate(key string, unit cacheUnit) {
	mx := k.getKeyMx(key)
	mx.Lock()
	go func(key string, u cacheUnit) {
		if time.Now().Unix()-u.tta <= TTL_CACHE_SEC {
			v, err := k.db.Get(key)
			k.addToCache(key, v, err)
		}
		mx.Unlock()
	}(key, unit)
}

func (k *KVCache) update(key string) (string, error) {
	mx := k.getKeyMx(key)

	mx.Lock()
	metric.Add(1)
	v, err := k.db.Get(key) //all db are waiting
	mx.Unlock()

	k.addToCache(key, v, err)

	return v, err
}

func (k *KVCache) addToCache(key string, value string, err error) {
	k.mu.Lock()
	k.m[key] = cacheUnit{v: value, err: err, tta: time.Now().Unix()}
	k.mu.Unlock()
}

func (k *KVCache) Set(key, value string) error {
	return k.db.Set(key, value)
}

type KVDB struct {
	mu sync.RWMutex
	m  map[string]string
}

func NewKVDB() *KVDB {
	return &KVDB{
		m: make(map[string]string),
	}
}

func (c *KVDB) Get(key string) (string, error) {
	c.mu.RLock()
	v, ok := c.m[key]
	c.mu.RUnlock()
	if ok {
		return v, nil
	}

	return "", KeyNotFound
}

func (c *KVDB) Set(key, value string) error {
	c.mu.Lock()
	c.m[key] = value
	c.mu.Unlock()

	return nil
}

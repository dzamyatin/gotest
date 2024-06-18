package lib

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

//
//func main() {
//	wg := sync.WaitGroup{}
//
//	db := NewKVDB()
//	cache := NewKVCache(db)
//
//	for i := 0; i < 100; i++ {
//		go func(i int) {
//			cache.Set("Hella", strconv.Itoa(i))
//		}(i)
//	}
//
//	strt := time.Now()
//	for i := 0; i < 1000000; i++ {
//		//for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go func() {
//			cache.Get("Hella")
//			//fmt.Println(
//			//	cache.Get("Hella"),
//			//)
//			wg.Done()
//		}()
//	}
//
//	wg.Wait()
//
//	log.Println("Met:")
//	log.Println(KvMetric.Load())
//	log.Println("Met2:")
//	log.Println(KvMetric2.Load())
//	end := time.Now().Sub(strt)
//	fmt.Printf("Time: %s", end)
//}

var KvMetric = atomic.Int32{}
var KvMetric2 = atomic.Int32{}

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
	if v, ok := k.getFromCache(key); ok {
		k.invalidate(key)
		return v.v, v.err
	}

	return k.update(key)
}

func (k *KVCache) isMx(key string) bool {
	k.mu.RLock()
	_, ok := k.m[key]
	k.mu.RUnlock()

	return ok
}

func (k *KVCache) getFromCache(key string) (cacheUnit, bool) {
	k.mu.RLock()

	if v, ok := k.m[key]; ok {
		k.mu.RUnlock()
		return v, true
	}

	k.mu.RUnlock()

	return cacheUnit{}, false
}

func (k *KVCache) getKeyMx(key string) *sync.RWMutex {
	k.mxx.Lock()
	defer k.mxx.Unlock()

	if v, ok := k.mx[key]; ok {
		return v
	}

	mx := &sync.RWMutex{}
	k.mx[key] = mx

	return mx
}

func (k *KVCache) deleteKeyMx(key string) {
	k.mxx.Lock()
	defer k.mxx.Unlock()

	delete(k.mx, key)
}

func (k *KVCache) invalidate(key string) {
	k.keyAtomic(
		key,
		func(key string, mx *sync.RWMutex) {
			KvMetric.Add(1)
			u, ok := k.getFromCache(key)

			if !ok || time.Now().Unix()-u.tta > TTL_CACHE_SEC {
				go func(key string) {
					v, err := k.db.Get(key)
					k.addToCache(key, v, err)

					mx.Unlock()
				}(key)
			} else {
				mx.Unlock()
			}
		},
	)
}

func (k *KVCache) keyAtomic(key string, f func(string, *sync.RWMutex)) {
	mx := k.getKeyMx(key)
	mx.Lock()
	f(key, mx)
	k.deleteKeyMx(key)
}

func (k *KVCache) update(key string) (string, error) {
	var v string
	var err error

	if !k.isMx(key) {
		k.keyAtomic(
			key,
			func(key string, mx *sync.RWMutex) {
				v, err = k.db.Get(key) //all db are waiting
				mx.Unlock()

				k.addToCache(key, v, err)
			},
		)
	} else {
		v, err = k.Get(key)
	}

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
	KvMetric2.Add(1)
	//time.Sleep(100 * time.Millisecond)

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

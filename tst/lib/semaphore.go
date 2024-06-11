package lib

import "sync"

//Task: semaphore

type SemaphoreChan struct {
	c chan struct{}
}

func NewSemaphoreChan(count int) SemaphoreChan {
	return SemaphoreChan{
		c: make(chan struct{}, count),
	}
}

func (s *SemaphoreChan) Lock() {
	s.c <- struct{}{}
}

func (s *SemaphoreChan) Release() {
	<-s.c
}

type SemaphoreSync struct {
	i  int
	mx sync.RWMutex
}

func NewSemaphoreSync(count int) SemaphoreSync {
	return SemaphoreSync{i: count}
}

func (s *SemaphoreSync) Lock() {
	for {
		s.mx.Lock()
		if 0 < s.i {
			break
		}
		s.mx.Unlock()
	}

	s.i--
	s.mx.Unlock()
}

func (s *SemaphoreSync) Release() {
	s.mx.Lock()
	s.i++
	s.mx.Unlock()
}

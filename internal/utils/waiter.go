package utils

import "sync"

type Waiter struct {
	done chan struct{}
	once sync.Once
}

func NewWaiter() *Waiter {
	return &Waiter{
		done: make(chan struct{}),
	}
}

func (s *Waiter) Wait() {
	<-s.done
}

func (s *Waiter) Broadcast() {
	s.once.Do(func() {
		close(s.done)
	})
}

package utils

import (
	"sync"
)

type Queue[T any] struct {
	mu   sync.Mutex
	cond *sync.Cond
	head *node[T]
	tail *node[T]
}

type node[T any] struct {
	value T
	next  *node[T]
}

func NewQueue[T any]() *Queue[T] {
	dummy := &node[T]{}
	q := &Queue[T]{
		head: dummy,
		tail: dummy,
	}
	q.cond = sync.NewCond(&q.mu)
	return q
}

func (q *Queue[T]) Enqueue(v T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	newNode := &node[T]{value: v}
	q.tail.next = newNode
	q.tail = newNode

	q.cond.Signal()
}

func (q *Queue[T]) Dequeue() T {
	q.mu.Lock()
	defer q.mu.Unlock()

	for q.head.next == nil {
		q.cond.Wait()
	}

	current := q.head.next
	v := current.value

	q.head.next = current.next

	if q.head.next == nil {
		q.tail = q.head
	}

	return v
}

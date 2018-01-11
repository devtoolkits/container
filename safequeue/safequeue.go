package safequeue

import (
	"container/list"
	"sync"
)

const UNLIMITED = -1

type Queue struct {
	sync.RWMutex
	container *list.List
	capacity  int
}

func NewQueue() *Queue {
	return NewCappedQueue(UNLIMITED)
}

func NewCappedQueue(cap int) *Queue {
	return &Queue{container: list.New(), capacity: cap}
}

// Push insert 1 element to queue
func (q *Queue) Push(item interface{}) bool {
	q.Lock()
	defer q.Unlock()

	if q.capacity == UNLIMITED || q.container.Len() < q.capacity {
		q.container.PushFront(item)
		return true
	}
	return false
}

// Pop return 1 element from queue
func (q *Queue) Pop() interface{} {
	q.Lock()
	defer q.Unlock()

	item := q.container.Back()
	if item != nil {
		return q.container.Remove(item)
	}
	return nil
}

// Pop return N(max) element from queue
func (q *Queue) PopN(n int) []interface{} {
	q.Lock()
	defer q.Unlock()

	ret := make([]interface{}, 0, n)
	for i := 0; i < n; i++ {
		temp := q.container.Back()
		if temp == nil {
			break
		}
		ret = append(ret, temp.Value)
		q.container.Remove(temp)
	}

	return ret
}

// Size
func (q *Queue) Size() int {
	q.RLock()
	defer q.RUnlock()

	return q.container.Len()
}

// Full is queue full?
func (q *Queue) Full() bool {
	q.RLock()
	defer q.RUnlock()

	if q.capacity == UNLIMITED {
		return false
	}
	return q.container.Len() >= q.capacity
}

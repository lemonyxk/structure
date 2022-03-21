package queue

import (
	"sync"
)

func NewBlockQueue[T any]() *blockQueue[T] {
	var queue = &blockQueue[T]{}
	queue.cond = sync.NewCond(new(sync.RWMutex))
	queue.status = blockQueueStatus{wait: 0, len: 0}
	return queue
}

type blockQueueStatus struct {
	wait int
	len  int
}

type blockQueue[T any] struct {
	cond    *sync.Cond
	storage []T
	status  blockQueueStatus
}

func (queue *blockQueue[T]) Push(v T) {

	queue.cond.L.Lock()

	queue.storage = append(queue.storage, v)
	queue.status.len++

	if queue.status.wait > 0 {
		queue.cond.Signal()
	}

	queue.cond.L.Unlock()
}

func (queue *blockQueue[T]) Pop() T {

	queue.cond.L.Lock()

	queue.status.wait++

	for {
		if len(queue.storage) > 0 {
			var r = queue.storage[0]
			queue.storage = queue.storage[1:]
			queue.status.wait--
			queue.status.len--
			queue.cond.L.Unlock()
			return r
		}
		queue.cond.Wait()
	}
}

func (queue *blockQueue[T]) Status() blockQueueStatus {
	return queue.status
}

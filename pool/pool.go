package pool

import (
	"runtime"
	"sync"
)

type Config[T any] struct {
	Max int
	Min int
	New func() T
}

func New[T any](config Config[T]) *Pool[T] {

	if config.Max <= 0 {
		config.Max = runtime.NumCPU() * 2
	}

	if config.Min <= 0 {
		config.Min = runtime.NumCPU()
	}

	if config.Min >= config.Max {
		config.Min = config.Max
	}

	var pool = &Pool[T]{}
	pool.config = config
	pool.status = Status{Max: config.Max, Min: config.Min, Len: 0}
	//
	// if len(pool.storage) < pool.config.Min {
	// 	pool.storage = append(pool.storage, config.New())
	// }

	return pool
}

type Status struct {
	Max int
	Min int
	Len int
}

type Pool[T any] struct {
	mux     sync.RWMutex
	storage []T
	config  Config[T]
	status  Status
}

func (pool *Pool[T]) Put(v T) {

	pool.mux.Lock()
	defer pool.mux.Unlock()

	if len(pool.storage) >= pool.config.Max {
		return
	}

	// if put too fast and get slowly that you will lose some put things
	// pool not need worry
	if len(pool.storage) < pool.config.Max {
		pool.storage = append(pool.storage, v)
		pool.status.Len++
	}
}

func (pool *Pool[T]) Get() T {

	pool.mux.Lock()
	defer pool.mux.Unlock()

	if len(pool.storage) > 0 {
		var r = pool.storage[0]
		pool.storage = pool.storage[1:]
		pool.status.Len--
		return r
	}

	return pool.config.New()
}

func (pool *Pool[T]) Status() Status {
	return pool.status
}

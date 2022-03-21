package pool

import (
	"runtime"
	"sync"
)

type LastPoolConfig[T any] struct {
	Max int
	Min int
	New func() T
}

func NewLastPool[T any](config LastPoolConfig[T]) *lastPool[T] {

	if config.Max <= 0 {
		config.Max = runtime.NumCPU() * 2
	}

	if config.Min <= 0 {
		config.Min = runtime.NumCPU()
	}

	if config.Min >= config.Max {
		config.Min = config.Max
	}

	var pool = &lastPool[T]{}
	pool.config = config
	pool.status = lastPoolStatus{max: config.Max, min: config.Min, len: 0}

	if len(pool.storage) < pool.config.Min {
		pool.storage = append(pool.storage, config.New())
	}

	return pool
}

type lastPoolStatus struct {
	max int
	min int
	len int
}

type lastPool[T any] struct {
	mux     sync.RWMutex
	storage []T
	config  LastPoolConfig[T]
	status  lastPoolStatus
}

func (pool *lastPool[T]) Put(v T) {

	if len(pool.storage) >= pool.config.Max {
		return
	}

	pool.mux.Lock()

	// if put too fast and get slowly that you will lose some put things
	// pool do not need worry
	if len(pool.storage) < pool.config.Max {
		pool.storage = append(pool.storage, v)
		pool.status.len++
	}

	pool.mux.Unlock()
}

func (pool *lastPool[T]) Get() T {

	pool.mux.Lock()

	if len(pool.storage) > 0 {
		var r = pool.storage[0]
		pool.storage = pool.storage[1:]
		pool.status.len--
		pool.mux.Unlock()
		return r
	}

	pool.mux.Unlock()
	return pool.config.New()

}

func (pool *lastPool[T]) Status() lastPoolStatus {
	return pool.status
}

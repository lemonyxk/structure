/**
* @program: structure
*
* @description:
*
* @author: lemo
*
* @create: 2022-03-21 23:38
**/

package hash

import "sync"

type Hash[K comparable, V any] struct {
	m   map[K]V
	mux sync.RWMutex
}

func New[K comparable, V any]() *Hash[K, V] {
	return &Hash[K, V]{
		m: make(map[K]V),
	}
}

func (h *Hash[K, V]) Set(k K, v V) {
	h.mux.Lock()
	h.m[k] = v
	h.mux.Unlock()
}

func (h *Hash[K, V]) Get(k K) V {
	h.mux.RLock()
	v := h.m[k]
	h.mux.RUnlock()
	return v
}

func (h *Hash[K, V]) Delete(k K) {
	h.mux.Lock()
	delete(h.m, k)
	h.mux.Unlock()
}

func (h *Hash[K, V]) Len() int {
	h.mux.RLock()
	var l = len(h.m)
	h.mux.RUnlock()
	return l
}

func (h *Hash[K, V]) Clear() {
	h.mux.Lock()
	h.m = make(map[K]V)
	h.mux.Unlock()
}

func (h *Hash[K, V]) Range(fn func(k K, v V) bool) {
	h.mux.RLock()
	for k, v := range h.m {
		if !fn(k, v) {
			break
		}
	}
	h.mux.RUnlock()
}

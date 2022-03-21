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
	mux sync.Mutex
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
	h.mux.Lock()
	v := h.m[k]
	h.mux.Unlock()
	return v
}

func (h *Hash[K, V]) Delete(k K) {
	h.mux.Lock()
	delete(h.m, k)
	h.mux.Unlock()
}

func (h *Hash[K, V]) Len() int {
	return len(h.m)
}

func (h *Hash[K, V]) Clear() {
	h.mux.Lock()
	h.m = make(map[K]V)
	h.mux.Unlock()
}

func (h *Hash[K, V]) Range(fn func(k K, v V) bool) {
	h.mux.Lock()
	defer h.mux.Unlock()
	for k, v := range h.m {
		h.mux.Unlock()
		if !fn(k, v) {
			break
		}
		h.mux.Lock()
	}
}

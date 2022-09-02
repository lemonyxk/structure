/**
* @program: structure
*
* @description:
*
* @author: lemo
*
* @create: 2022-03-22 01:53
**/

package hash

import (
	"sync"
)

type SyncHash[K, V any] struct {
	m   sync.Map
	len int
}

func NewSync[K, V any]() *SyncHash[K, V] {
	return &SyncHash[K, V]{}
}

func (s *SyncHash[K, V]) Set(key K, value V) {
	s.m.Store(key, value)
}

func (s *SyncHash[K, V]) Get(key K) V {
	var v, b = s.m.Load(key)
	if b {
		return v.(V)
	}
	var r V
	return r
}

func (s *SyncHash[K, V]) Delete(key K) {
	s.m.Delete(key)
}

func (s *SyncHash[K, V]) Range(fn func(key K, value V) bool) {
	s.m.Range(func(key any, value any) bool {
		return fn(key.(K), value.(V))
	})
}

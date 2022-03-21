/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2019-11-14 12:34
**/

package stack

type stack[T any] struct {
	list []T
}

func NewStack[T any](list ...T) *stack[T] {
	return &stack[T]{list: list}
}

func (s *stack[T]) Push(v T) {
	s.list = append(s.list, v)
}

func (s *stack[T]) Pop() (T, bool) {
	if len(s.list) == 0 {
		var t T
		return t, false
	}
	var v = s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return v, true
}

func (s *stack[T]) Top() (T, bool) {
	if len(s.list) == 0 {
		var t T
		return t, false
	}
	return s.list[len(s.list)-1], true
}

func (s *stack[T]) Size() int {
	return len(s.list)
}

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

type Stack[T any] struct {
	list []T
}

func New[T any](list ...T) *Stack[T] {
	return &Stack[T]{list: list}
}

func (s *Stack[T]) Push(v T) {
	s.list = append(s.list, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.list) == 0 {
		var t T
		return t, false
	}
	var v = s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return v, true
}

func (s *Stack[T]) Top() (T, bool) {
	if len(s.list) == 0 {
		var t T
		return t, false
	}
	return s.list[len(s.list)-1], true
}

func (s *Stack[T]) Size() int {
	return len(s.list)
}

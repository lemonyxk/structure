/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2019-11-14 12:14
**/

package queue

func New[T any](list ...T) *Queue[T] {
	return &Queue[T]{list: list}
}

type Queue[T any] struct {
	list []T
}

func (q *Queue[T]) Push(v T) {
	q.list = append(q.list, v)
}

func (q *Queue[T]) Pop() (T, bool) {
	if len(q.list) == 0 {
		var t T
		return t, false
	}
	var v = q.list[0]
	q.list = q.list[1:]
	return v, true
}

func (q *Queue[T]) Top() (T, bool) {
	if len(q.list) == 0 {
		var t T
		return t, false
	}
	return q.list[0], true
}

func (q *Queue[T]) Size() int {
	return len(q.list)
}

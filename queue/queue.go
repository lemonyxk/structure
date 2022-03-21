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

func New[T any](list ...T) *queue[T] {
	return &queue[T]{list: list}
}

type queue[T any] struct {
	list []T
}

func (q *queue[T]) Push(v T) {
	q.list = append(q.list, v)
}

func (q *queue[T]) Pop() (T, bool) {
	if len(q.list) == 0 {
		var t T
		return t, false
	}
	var v = q.list[0]
	q.list = q.list[1:]
	return v, true
}

func (q *queue[T]) Top() (T, bool) {
	if len(q.list) == 0 {
		var t T
		return t, false
	}
	return q.list[0], true
}

func (q *queue[T]) Size() int {
	return len(q.list)
}

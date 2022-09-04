/**
* @program: structure
*
* @description:
*
* @author: lemo
*
* @create: 2022-06-09 10:08
**/

package scheduler

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/lemonyxk/structure/v3/queue"
)

func New() *Scheduler {
	var s = &Scheduler{
		queue:  queue.NewBlock[*Worker](),
		status: Status{running: 0},
		limit:  0,
	}

	go s.loop()

	return s
}

type Status struct {
	running int32
}

type Scheduler struct {
	mux    sync.Mutex
	queue  *queue.Block[*Worker]
	status Status
	limit  int32
}

type Worker struct {
	fn      func()
	timeout <-chan time.Time
	stop    chan struct{}
	ch      chan error
}

func (w *Worker) Stop() {
	w.stop <- struct{}{}
}

func (w *Worker) Timeout(timeout time.Duration) *Worker {
	w.timeout = time.After(timeout)
	return w
}

func (w *Worker) Wait() error {
	return <-w.ch
}

func (s *Scheduler) Status() Status {
	return s.status
}

func (s *Scheduler) SetLimit(limit int32) {
	s.limit = limit
}

func (s *Scheduler) Add(fn func()) *Worker {
	w := &Worker{
		fn:      fn,
		timeout: nil,
		stop:    make(chan struct{}, 1),
		ch:      make(chan error, 1),
	}

	s.queue.Push(w)

	return w
}

func (s *Scheduler) loop() {
	for {

		if s.limit > 0 {
			for atomic.LoadInt32(&s.status.running) >= s.limit {
			}
		}

		var w = s.queue.Pop()

		atomic.AddInt32(&s.status.running, 1)

		var sucCh = make(chan struct{}, 1)

		go func() {
			w.fn()
			sucCh <- struct{}{}
		}()

		go func() {
			select {
			case <-w.timeout:
				w.ch <- errors.New("timeout")
			case <-w.stop:
				w.ch <- errors.New("stop")
			case <-sucCh:
				w.ch <- nil
			}
			atomic.AddInt32(&s.status.running, -1)
		}()
	}
}

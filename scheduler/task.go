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
}

type Worker struct {
	fn      func(chan struct{}, chan error)
	timeout <-chan time.Time
	stop    chan struct{}
}

func (w *Worker) Stop() {
	w.stop <- struct{}{}
}

func (s *Scheduler) Status() Status {
	return s.status
}

func (s *Scheduler) Add(fn func(chan struct{}, chan error), timeout time.Duration) *Worker {
	w := &Worker{
		fn:      fn,
		timeout: time.After(timeout),
		stop:    make(chan struct{}),
	}

	s.queue.Push(w)

	return w
}

func (s *Scheduler) loop() {
	for {

		for atomic.LoadInt32(&s.status.running) > 10 {
		}

		var w = s.queue.Pop()

		atomic.AddInt32(&s.status.running, 1)

		var sucCh = make(chan struct{}, 1)
		var errCh = make(chan error, 1)

		go func() {
			w.fn(sucCh, errCh)
		}()

		go func() {
			select {
			case <-w.timeout:
				errCh <- errors.New("timeout")
			case <-w.stop:
				errCh <- errors.New("stop")
			case <-sucCh:
			}
			atomic.AddInt32(&s.status.running, -1)
		}()
	}
}

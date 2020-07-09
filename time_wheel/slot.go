/**
* @program: timewheel
*
* @description:
*
* @author: lemo
*
* @create: 2020-07-08 02:33
**/

package timewheel

import (
	"sync"
	"time"
)

type slot struct {
	data []*Task
	mux  sync.RWMutex
}

func (n *slot) addTask(times int, cal calibration, duration time.Duration, fn func()) *Task {
	n.mux.Lock()
	defer n.mux.Unlock()
	var task = &Task{times: times, cal: cal, duration: duration, fn: fn}
	n.data = append(n.data, task)
	return task
}

func (n *slot) setTask(task *Task) {
	n.mux.Lock()
	defer n.mux.Unlock()
	n.data = append(n.data, task)
}

func (n *slot) getAndReset() []*Task {
	n.mux.Lock()
	defer n.mux.Unlock()
	var res = n.data
	n.data = n.data[0:0]
	return res
}

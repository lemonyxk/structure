/**
* @program: timewheel
*
* @description:
*
* @author: lemo
*
* @create: 2020-07-08 02:29
**/

package timewheel

import (
	"sync"
	"time"
)

var acc = 100 // ms
var msLen = 1000 / acc

const secLen = 60
const minLen = 60
const hourLen = 24

type TimeWheel struct {
	ms   []*slot
	sec  []*slot
	min  []*slot
	hour []*slot

	cal       calibration
	mux       sync.Mutex
	tMux      sync.Mutex
	isRunning bool
}

func NewTimeWheel(accuracy int) *TimeWheel {
	// suggest 10 or 100 ms
	if accuracy < 0 || accuracy > 1000 {
		panic("accuracy must between 1-1000 ms")
	}
	acc = accuracy
	msLen = 1000 / acc
	var timeWheel = &TimeWheel{
		ms:   initSlots(msLen),
		sec:  initSlots(secLen),
		min:  initSlots(minLen),
		hour: initSlots(hourLen),
	}
	return timeWheel
}

func initSlots(slotLen int) []*slot {
	var slots []*slot
	for i := 0; i < slotLen; i++ {
		var slot = &slot{}
		// slot.init()
		slots = append(slots, slot)
	}
	return slots
}

func getPosition(duration time.Duration) calibration {

	var mill = int(duration.Milliseconds())

	if mill == 0 {
		mill = acc
	} else {
		if mill%acc != 0 {
			mill += acc
		}
	}

	var all = mill / acc

	var ms = 0
	var sec = 0
	var min = 0
	var hour = 0

	ms = all % msLen
	all /= msLen

	sec = all % secLen
	all /= secLen

	min = all % minLen
	all /= minLen

	hour = all % hourLen
	all /= hourLen

	return calibration{ms: ms, sec: sec, min: min, hour: hour}
}

func (tw *TimeWheel) getSlot(cal calibration) *slot {
	if cal.ms != 0 {
		return tw.ms[cal.ms]
	}

	if cal.sec != 0 {
		return tw.sec[cal.sec]
	}

	if cal.min != 0 {
		return tw.min[cal.min]
	}

	if cal.hour != 0 {
		return tw.hour[cal.hour]
	}

	return tw.hour[cal.hour]
}

func (tw *TimeWheel) Stop() {
	tw.tMux.Lock()
	defer tw.tMux.Unlock()

	if !tw.isRunning {
		return
	}

	tw.isRunning = false
}

func (tw *TimeWheel) Start() *TimeWheel {

	tw.tMux.Lock()
	defer tw.tMux.Unlock()

	if tw.isRunning {
		return nil
	}

	tw.isRunning = true

	var ticker = time.NewTicker(time.Millisecond * time.Duration(acc))

	go func() {
		for range ticker.C {
			var change = tw.setCalibration(calibration{ms: 1})

			// ms task
			if change&msChange != 0 {
				tw.runTask(tw.ms[tw.cal.ms])
			}
			// sec task
			if change&secChange != 0 {
				tw.runTask(tw.sec[tw.cal.sec])
			}
			// min task
			if change&minChange != 0 {
				tw.runTask(tw.min[tw.cal.min])
			}
			// hour task
			if change&hourChange != 0 {
				tw.runTask(tw.hour[tw.cal.hour])
			}

			if !tw.isRunning {
				break
			}
		}
	}()

	return tw
}

func (tw *TimeWheel) runTask(slot *slot) {

	var nodes = slot.getAndReset()

	for i := 0; i < len(nodes); i++ {
		// has delete
		if nodes[i].hasDelete {
			continue
		}
		// equal
		if !equal(tw.cal, nodes[i].cal) {
			// not the time
			// put in again
			slot.setTask(nodes[i])
			continue
		}
		// limit times
		if nodes[i].times > 0 {
			nodes[i].times--
			if nodes[i].times > 0 {
				var cal = getPosition(nodes[i].duration)
				var merge, _ = merge(nodes[i].cal, cal)
				nodes[i].cal = merge
				tw.getSlot(merge).setTask(nodes[i])
			}
			go nodes[i].fn()
			continue
		}
		// unlimited times
		if nodes[i].times < 0 {
			var cal = getPosition(nodes[i].duration)
			var merge, _ = merge(nodes[i].cal, cal)
			nodes[i].cal = merge
			tw.getSlot(merge).setTask(nodes[i])
			go nodes[i].fn()
			continue
		}
	}

}

func (tw *TimeWheel) setCalibration(c1 calibration) int {
	tw.mux.Lock()
	defer tw.mux.Unlock()
	var cal, change = merge(tw.cal, c1)
	tw.cal = cal
	return change
}

func (tw *TimeWheel) Timer(duration time.Duration, fn func()) *Task {
	return tw.addTask(1, duration, fn)
}

func (tw *TimeWheel) Ticker(duration time.Duration, fn func()) *Task {
	return tw.addTask(-1, duration, fn)
}

func (tw *TimeWheel) Times(duration time.Duration, times int, fn func()) *Task {
	if times < 1 {
		panic("times at least 1")
	}
	return tw.addTask(times, duration, fn)
}

func (tw *TimeWheel) addTask(times int, duration time.Duration, fn func()) *Task {
	if duration.Milliseconds() > time.Second.Milliseconds()*86400 {
		panic("more than one day")
	}

	// put task at next round
	tw.mux.Lock()
	defer tw.mux.Unlock()

	var cal = getPosition(duration)
	var merge, _ = merge(tw.cal, cal)
	return tw.getSlot(merge).addTask(times, merge, duration, fn)
}

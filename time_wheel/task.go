/**
* @program: timewheel
*
* @description:
*
* @author: lemo
*
* @create: 2020-07-09 16:30
**/

package timewheel

import "time"

type Task struct {
	hasDelete bool
	cal       calibration
	fn        func()
	times     int
	duration  time.Duration
}

func (t *Task) Delete() {
	t.hasDelete = true
}

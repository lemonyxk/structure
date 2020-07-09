/**
* @program: timewheel
*
* @description:
*
* @author: lemo
*
* @create: 2020-07-08 09:42
**/

package timewheel

import "fmt"

const (
	msChange = 1 << iota
	secChange
	minChange
	hourChange
)

type calibration struct {
	ms   int
	sec  int
	min  int
	hour int
}

func (cal calibration) String() string {
	return fmt.Sprintf("%d Hour %d Min %d Sec %d Ms", cal.hour, cal.min, cal.sec, cal.ms)
}

func equal(c1 calibration, c2 calibration) bool {
	return c1.ms == c2.ms && c1.sec == c2.sec && c1.min == c2.min && c1.hour == c2.hour
}

func merge(c1 calibration, c2 calibration) (calibration, int) {
	var change = 0

	if c2.ms != 0 {
		change |= msChange
	}

	c1.ms += c2.ms
	if c1.ms >= msLen {
		c1.ms %= msLen
		c1.sec++
		change |= secChange
	}
	c1.sec += c2.sec
	if c1.sec >= secLen {
		c1.sec %= secLen
		c1.min++
		change |= minChange
	}
	c1.min += c2.min
	if c1.min >= minLen {
		c1.min %= minLen
		c1.hour++
		change |= hourChange
	}
	c1.hour += c2.hour
	if c1.hour >= hourLen {
		c1.hour %= hourLen
	}

	return c1, change
}

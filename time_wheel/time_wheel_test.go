/**
* @program: timewheel
*
* @description:
*
* @author: lemo
*
* @create: 2020-07-09 03:56
**/

package timewheel

import (
	"testing"
	"time"
)

var timeWheel = NewTimeWheel(10).Start()

// for i := 0; i < 100000; i++ {
// 	time.AfterFunc(time.Millisecond * 10 * time.Duration(i+1), func() {
// 		log.Println(time.Now())
// 	})
// }

func BenchmarkNewTimeWheel(b *testing.B) {
	for j := 0; j <= b.N; j++ {
		timeWheel.Ticker(time.Millisecond*10*time.Duration(j+1), func() {
			// log.Println(time.Now())
		})
	}
}

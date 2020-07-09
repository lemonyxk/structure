/**
* @program: timewheel
*
* @description:
*
* @author: lemo
*
* @create: 2020-07-08 02:28
**/

package main

func main() {

	// mac pro 2017 cpu: 50-60%
	// var timeWheel = timewheel.NewTimeWheel(10).Start()
	// for i := 0; i < 8640000; i++ {
	// 	timeWheel.Ticker(time.Millisecond*10*time.Duration(i+1), func() {
	// 		// log.Println(a+1, time.Now())
	// 	})
	// }

	// mac pro 2017 cpu: 400+%
	// for i := 0; i < 8640000; i++ {
	// 	var ticker = time.NewTicker(time.Millisecond * 10 * time.Duration(i+1))
	// 	go func() {
	// 		for range ticker.C {
	//
	// 		}
	// 	}()
	// }

	select {}

	// log.Println(getPosition(time.Second * 86400))

}

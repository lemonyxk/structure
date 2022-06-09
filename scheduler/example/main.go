/**
* @program: structure
*
* @description:
*
* @author: lemo
*
* @create: 2022-06-09 10:10
**/

package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/lemonyxk/structure/v3/scheduler"
)

func main() {
	// http.HandleFunc("/runTask", runTask)
	// http.HandleFunc("/stopTask", stopTask)
	//
	// go http.ListenAndServe(":9999", nil)

	var wait sync.WaitGroup

	wait.Add(1000)

	var counter int32 = 0

	var s = scheduler.New()

	for i := 1; i <= 1000; i++ {

		var index = i

		go func() {

			var f1 = func(sucCh chan struct{}, errCh chan error) {

				defer wait.Done()

				time.Sleep(time.Millisecond * 10)

				select {
				case err := <-errCh:
					fmt.Println(err, index)
					return
				default:
				}

				// time.Sleep(time.Second * 3)

				fmt.Println("task:", index)
				atomic.AddInt32(&counter, int32(index))
				sucCh <- struct{}{}
				// log.Println("sucCh done")
			}

			var worker = s.Add(f1, time.Millisecond*1000)

			_ = worker

			time.AfterFunc(time.Millisecond*1, func() {
				// worker.Stop()
			})

		}()
	}

	go func() {

		for {
			time.Sleep(time.Second * 1)
			log.Println("running:", s.Status(), "counter:", counter)
		}

	}()

	wait.Wait()

	log.Println(counter)
}

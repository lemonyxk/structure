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
	"log"
	"sync"
	"time"

	"github.com/lemonyxk/structure/scheduler"
)

func main() {
	// http.HandleFunc("/runTask", runTask)
	// http.HandleFunc("/stopTask", stopTask)
	//
	// go http.ListenAndServe(":9999", nil)

	// var wait sync.WaitGroup
	//
	// wait.Add(1000)
	//
	// var counter int32 = 0
	//
	// var s = scheduler.New()
	//
	// s.SetLimit(10)
	//
	// for i := 1; i <= 1000; i++ {
	//
	// 	var index = i
	//
	// 	go func() {
	//
	// 		var f1 = func() {
	//
	// 			time.Sleep(time.Millisecond * 10)
	//
	// 			// time.Sleep(time.Second * 3)
	//
	// 			// fmt.Println("task:", index)
	// 			// log.Println("sucCh done")
	// 		}
	//
	// 		var worker = s.Add(f1).Timeout(time.Millisecond * 15)
	//
	// 		time.AfterFunc(time.Millisecond*1, func() {
	// 			// worker.Stop()
	// 		})
	//
	// 		var err = worker.Wait()
	// 		if err != nil {
	// 			fmt.Println(err, index)
	// 		} else {
	// 			fmt.Println("task done", index)
	// 			atomic.AddInt32(&counter, int32(index))
	// 		}
	//
	// 		wait.Done()
	// 	}()
	// }
	//
	// go func() {
	//
	// 	for {
	// 		time.Sleep(time.Second * 1)
	// 		log.Println("running:", s.Status(), "counter:", counter)
	// 	}
	//
	// }()
	//
	// wait.Wait()
	//
	// log.Println(counter)

	var wait sync.WaitGroup

	wait.Add(10)

	var manager = scheduler.New()

	manager.SetLimit(1)

	for i := 0; i < 10; i++ {

		var index = i

		go func() {

			var f1 = func() {
				time.Sleep(time.Second * 1)
				log.Println("hello task", index)
			}

			var worker = manager.Add(f1)

			if index == 5 {
				time.AfterFunc(time.Millisecond*500, func() {
					worker.Stop()
				})
			}

			var err = worker.Wait()
			if err != nil {
				log.Println("err:", err)
			}

			wait.Done()
		}()
	}

	wait.Wait()
}

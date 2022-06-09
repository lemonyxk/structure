/**
* @program: structure
*
* @description:
*
* @author: lemo
*
* @create: 2022-03-20 03:23
**/

package main

import (
	"log"

	"github.com/lemonyxk/structure/v3/queue"
)

func main() {
	var p = queue.NewBlock[int]()

	p.Push(1)

	var a = p.Pop()

	var b = 1

	b = a

	log.Println(a, b)

	var p1 = queue.New(6, 1, 2, 3, 4, 5)

	for {
		v, ok := p1.Pop()
		if !ok {
			break
		}
		log.Println(v)
	}

}

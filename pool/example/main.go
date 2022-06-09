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

	"github.com/lemonyxk/structure/v3/pool"
)

func main() {
	var p = pool.New(pool.Config[any]{
		Max: 0,
		Min: 0,
		New: func() any {
			return 0
		},
	})

	var a = p.Get()

	var b = "hello world"

	a = b

	log.Println(a, b)

}

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

import "github.com/lemonyxk/structure/stack"

func main() {

	var p = stack.New(6, 7, 8, 9, 10, 1, 2, 3, 4, 5)

	for {
		var v, ok = p.Pop()
		if !ok {
			break
		}
		println(v)
	}

}

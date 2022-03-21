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

import "github.com/lemoyxk/structure/v3/stack"

func main() {

	var p = stack.NewStack(6, 7, 8, 9, 10, 1, 2, 3, 4, 5)

	for {
		var v, ok = p.Pop()
		if !ok {
			break
		}
		println(v)
	}

}

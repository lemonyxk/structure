/**
* @program: structure
*
* @description:
*
* @author: lemo
*
* @create: 2022-03-20 03:09
**/

package main

import "github.com/lemoyxk/structure/v3/head"

func main() {

	var h = head.NewMaxHead(9, 15, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	var s = h.Size()

	for i := 0; i < s; i++ {
		println(h.Pop())
	}

	var h1 = head.NewMaxHead(9, 15, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	var s1 = h1.Size()

	for i := 0; i < s1; i++ {
		println(h1.Pop())
	}

}
